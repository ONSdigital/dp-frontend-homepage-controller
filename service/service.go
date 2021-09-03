package service

import (
	"context"
	"github.com/ONSdigital/dp-api-clients-go/v2/health"
	"github.com/ONSdigital/dp-api-clients-go/v2/image"
	"github.com/ONSdigital/dp-api-clients-go/v2/renderer"
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/clients/release_calendar"
	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	"github.com/ONSdigital/dp-frontend-homepage-controller/homepage"
	"github.com/ONSdigital/dp-frontend-homepage-controller/routes"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"strings"
)

// Service contains all the configs, server and clients to run the frontend homepage controller
type Service struct {
	Config             *config.Config
	routerHealthClient *health.Client
	HealthCheck        HealthChecker
	Server             HTTPServer
	clients            *homepage.Clients
	ServiceList        *ExternalServiceList
	HomePageClient     homepage.HomepageClienter
}

// Run the service
func Run(ctx context.Context, cfg *config.Config, serviceList *ExternalServiceList, buildTime, gitCommit, version string, svcErrors chan error) (svc *Service, err error) {
	log.Info(ctx, "running service")

	// Initialise Service struct
	svc = &Service{
		Config:      cfg,
		ServiceList: serviceList,
	}

	// Get health client for api router
	svc.routerHealthClient = serviceList.GetHealthClient("api-router", cfg.APIRouterURL)

	// Initialise clients
	svc.clients = &homepage.Clients{
		Renderer: renderer.New(cfg.RendererURL),
		Zebedee:  zebedee.NewWithHealthClient(svc.routerHealthClient),
		Babbage:  release_calendar.New(cfg.BabbageURL),
		ImageAPI: image.NewWithHealthClient(svc.routerHealthClient),
	}

	// Get healthcheck with checkers
	svc.HealthCheck, err = serviceList.GetHealthCheck(cfg, buildTime, gitCommit, version)
	if err != nil {
		log.Fatal(ctx, "failed to create health check", err)
		return nil, err
	}
	if err := svc.registerCheckers(ctx, cfg); err != nil {
		return nil, errors.Wrap(err, "unable to register checkers")
	}

	if cfg.IsPublishingMode {
		svc.HomePageClient = homepage.NewHomePagePublishingClient(svc.clients)
	} else {
		languages := strings.Split(cfg.Languages, ",")
		svc.HomePageClient = homepage.NewHomePageWebClient(svc.clients, cfg.CacheUpdateInterval, languages)
		go svc.HomePageClient.StartBackgroundUpdate(ctx, svcErrors)
	}

	// Initialise router
	r := mux.NewRouter()
	routes.Init(ctx, r, svc.HealthCheck.Handler, svc.HomePageClient)
	svc.Server = serviceList.GetHTTPServer(cfg.BindAddr, r)

	// Start Healthcheck and HTTP Server
	log.Info(ctx, "Starting server", log.Data{"config": cfg})
	svc.HealthCheck.Start(ctx)
	go func() {
		if err := svc.Server.ListenAndServe(); err != nil {
			svcErrors <- errors.Wrap(err, "failure in http listen and serve")
		}
	}()

	return svc, nil
}

// Close gracefully shuts the service down in the required order, with timeout
func (svc *Service) Close(ctx context.Context) error {
	timeout := svc.Config.GracefulShutdownTimeout
	log.Info(ctx, "commencing graceful shutdown", log.Data{"graceful_shutdown_timeout": timeout})
	ctx, cancel := context.WithTimeout(ctx, timeout)
	hasShutdownError := false

	go func() {
		defer cancel()

		// stop healthcheck, as it depends on everything else
		if svc.ServiceList.HealthCheck {
			svc.HealthCheck.Stop()
		}

		if svc.HomePageClient != nil {
			svc.HomePageClient.Close()
		}

		// stop any incoming requests
		if err := svc.Server.Shutdown(ctx); err != nil {
			log.Error(ctx, "failed to shutdown http server", err)
			hasShutdownError = true
		}
	}()

	// wait for shutdown success (via cancel) or failure (timeout)
	<-ctx.Done()

	// timeout expired
	if ctx.Err() == context.DeadlineExceeded {
		log.Error(ctx, "shutdown timed out", ctx.Err())
		return ctx.Err()
	}

	// other error
	if hasShutdownError {
		err := errors.New("failed to shutdown gracefully")
		log.Error(ctx, "failed to shutdown gracefully ", err)
		return err
	}

	log.Info(ctx, "graceful shutdown was successful")
	return nil
}

func (svc *Service) registerCheckers(ctx context.Context, cfg *config.Config) (err error) {

	hasErrors := false

	if err = svc.HealthCheck.AddCheck("frontend renderer", svc.clients.Renderer.Checker); err != nil {
		hasErrors = true
		log.Error(ctx, "failed to add frontend renderer checker", err)
	}

	if err = svc.HealthCheck.AddCheck("Babbage", svc.clients.Babbage.Checker); err != nil {
		hasErrors = true
		log.Error(ctx, "failed to add babbage checker", err)
	}

	if err = svc.HealthCheck.AddCheck("API router", svc.routerHealthClient.Checker); err != nil {
		hasErrors = true
		log.Error(ctx, "failed to add api router checker", err)
	}

	if hasErrors {
		return errors.New("Error(s) registering checkers for healthcheck")
	}
	return nil
}
