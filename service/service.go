package service

import (
	"context"

	"github.com/ONSdigital/dp-api-clients-go/health"
	"github.com/ONSdigital/dp-api-clients-go/image"
	"github.com/ONSdigital/dp-api-clients-go/renderer"
	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/clients/release_calendar"
	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	"github.com/ONSdigital/dp-frontend-homepage-controller/routes"
	"github.com/ONSdigital/dp-net/handlers"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/pkg/errors"
)

// Service contains all the configs, server and clients to run the frontend homepage controller
type Service struct {
	Config             *config.Config
	routerHealthClient *health.Client
	HealthCheck        HealthChecker
	Server             HTTPServer
	clients            *routes.Clients
	ServiceList        *ExternalServiceList
}

// Run the service
func Run(ctx context.Context, cfg *config.Config, serviceList *ExternalServiceList, buildTime, gitCommit, version string, svcErrors chan error) (svc *Service, err error) {
	log.Event(ctx, "running service", log.INFO)

	// Initialise Service struct
	svc = &Service{
		Config:      cfg,
		ServiceList: serviceList,
	}

	// Get health client for api router
	svc.routerHealthClient = serviceList.GetHealthClient("api-router", cfg.APIRouterURL)

	// Initialise clients
	svc.clients = &routes.Clients{
		Renderer: renderer.New(cfg.RendererURL),
		Zebedee:  zebedee.New(cfg.ZebedeeURL),
		Babbage:  release_calendar.New(cfg.BabbageURL),
		ImageAPI: image.NewWithHealthClient(svc.routerHealthClient),
	}

	// Get healthcheck with checkers
	svc.HealthCheck, err = serviceList.GetHealthCheck(cfg, buildTime, gitCommit, version)
	if err != nil {
		log.Event(ctx, "failed to create health check", log.FATAL, log.Error(err))
		return nil, err
	}
	if err := svc.registerCheckers(ctx, cfg); err != nil {
		return nil, errors.Wrap(err, "unable to register checkers")
	}

	// Initialise router
	r := mux.NewRouter()
	routes.Init(ctx, r, svc.HealthCheck.Handler, svc.clients)
	m := createMiddleware(cfg)
	svc.Server = serviceList.GetHTTPServer(cfg.BindAddr, m.Then(r))

	// Start Healthcheck and HTTP Server
	log.Event(ctx, "Starting server", log.Data{"config": cfg})
	svc.HealthCheck.Start(ctx)
	go func() {
		if err := svc.Server.ListenAndServe(); err != nil {
			svcErrors <- errors.Wrap(err, "failure in http listen and serve")
		}
	}()

	return svc, nil
}

// CreateMiddleware creates an Alice middleware chain of handlers
// to forward collectionID from cookie and locale from header
func createMiddleware(cfg *config.Config) alice.Chain {
	m := alice.New(handlers.CheckCookie(handlers.CollectionID))
	return m.Append(handlers.CheckHeader(handlers.Locale))
}

// Close gracefully shuts the service down in the required order, with timeout
func (svc *Service) Close(ctx context.Context) error {
	timeout := svc.Config.GracefulShutdownTimeout
	log.Event(ctx, "commencing graceful shutdown", log.Data{"graceful_shutdown_timeout": timeout}, log.INFO)
	ctx, cancel := context.WithTimeout(ctx, timeout)
	hasShutdownError := false

	go func() {
		defer cancel()

		// stop healthcheck, as it depends on everything else
		if svc.ServiceList.HealthCheck {
			svc.HealthCheck.Stop()
		}

		// stop any incoming requests
		if err := svc.Server.Shutdown(ctx); err != nil {
			log.Event(ctx, "failed to shutdown http server", log.Error(err), log.ERROR)
			hasShutdownError = true
		}
	}()

	// wait for shutdown success (via cancel) or failure (timeout)
	<-ctx.Done()

	// timeout expired
	if ctx.Err() == context.DeadlineExceeded {
		log.Event(ctx, "shutdown timed out", log.ERROR, log.Error(ctx.Err()))
		return ctx.Err()
	}

	// other error
	if hasShutdownError {
		err := errors.New("failed to shutdown gracefully")
		log.Event(ctx, "failed to shutdown gracefully ", log.ERROR, log.Error(err))
		return err
	}

	log.Event(ctx, "graceful shutdown was successful", log.INFO)
	return nil
}

func (svc *Service) registerCheckers(ctx context.Context, cfg *config.Config) (err error) {

	hasErrors := false

	if err = svc.HealthCheck.AddCheck("frontend renderer", svc.clients.Renderer.Checker); err != nil {
		hasErrors = true
		log.Event(ctx, "failed to add frontend renderer checker", log.Error(err))
	}

	if err = svc.HealthCheck.AddCheck("Zebedee", svc.clients.Zebedee.Checker); err != nil {
		hasErrors = true
		log.Event(ctx, "failed to add zebedee checker", log.Error(err))
	}

	if err = svc.HealthCheck.AddCheck("Babbage", svc.clients.Babbage.Checker); err != nil {
		hasErrors = true
		log.Event(ctx, "failed to add babbage checker", log.Error(err))
	}

	if err = svc.HealthCheck.AddCheck("API Router (to access Image API)", svc.routerHealthClient.Checker); err != nil {
		hasErrors = true
		log.Event(ctx, "failed to add image api checker", log.Error(err))
	}

	if hasErrors {
		return errors.New("Error(s) registering checkers for healthcheck")
	}
	return nil
}
