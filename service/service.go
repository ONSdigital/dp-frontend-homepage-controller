package service

import (
	"context"

	"github.com/ONSdigital/dp-api-clients-go/v2/health"
	//nolint:typecheck // assets may not exist as they are auto generated
	"github.com/ONSdigital/dp-frontend-homepage-controller/assets"
	"github.com/ONSdigital/dp-frontend-homepage-controller/cache"
	cachePrivate "github.com/ONSdigital/dp-frontend-homepage-controller/cache/private"
	cachePublic "github.com/ONSdigital/dp-frontend-homepage-controller/cache/public"
	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	"github.com/ONSdigital/dp-frontend-homepage-controller/homepage"
	"github.com/ONSdigital/dp-frontend-homepage-controller/routes"
	render "github.com/ONSdigital/dp-renderer/v2"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// Service contains all the configs, server and clients to run the frontend homepage controller
type Service struct {
	Cache              cache.List
	Config             *config.Config
	RouterHealthClient *health.Client
	HealthCheck        HealthChecker
	Server             HTTPServer
	Clients            *homepage.Clients
	ServiceList        *ExternalServiceList
	HomePageClient     homepage.Clienter
	RendererClient     homepage.RenderClient
}

func New() *Service {
	return &Service{}
}

// Run the service
func (svc *Service) Run(ctx context.Context, cfg *config.Config, serviceList *ExternalServiceList, svcErrors chan error) (err error) {
	log.Info(ctx, "Initialising service")

	// Initialise render client
	rend := render.NewWithDefaultClient(assets.Asset, assets.AssetNames, cfg.PatternLibraryAssetsPath, cfg.SiteDomain)

	// Initialise router
	r := mux.NewRouter()
	routes.Init(
		ctx,
		r,
		cfg,
		svc.Cache,
		svc.HealthCheck.Handler,
		svc.HomePageClient,
		rend,
	)

	svc.Server = serviceList.GetHTTPServer(cfg.BindAddr, r)

	// Start Healthcheck and HTTP Server
	svc.HealthCheck.Start(ctx)
	go func() {
		if err := svc.Server.ListenAndServe(); err != nil {
			svcErrors <- errors.Wrap(err, "failure in http listen and serve")
		}
	}()

	return nil
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

func (svc *Service) RegisterCheckers(ctx context.Context) (err error) {
	hasErrors := false
	if err = svc.HealthCheck.AddCheck("API router", svc.RouterHealthClient.Checker); err != nil {
		hasErrors = true
		log.Error(ctx, "failed to add api router checker", err)
	}
	if hasErrors {
		return errors.New("Error(s) registering checkers for healthcheck")
	}
	return nil
}

func (svc *Service) InitiateServiceList(cfg *config.Config, svcList *ExternalServiceList) {
	svc.Config = cfg
	svc.ServiceList = svcList
	svc.RouterHealthClient = svcList.GetHealthClient("api-router", cfg.APIRouterURL)
}

func (svc *Service) Init(ctx context.Context, cfg *config.Config, serviceList *ExternalServiceList, buildTime, gitCommit, version string, svcErrors chan error) (err error) {
	// Get healthcheck with checkers
	svc.HealthCheck, err = serviceList.GetHealthCheck(cfg, buildTime, gitCommit, version)
	if err != nil {
		log.Fatal(ctx, "failed to create health check", err)
		return err
	}
	// Initialise clients
	if registerErr := svc.RegisterCheckers(ctx); registerErr != nil {
		return errors.Wrap(registerErr, "unable to register checkers")
	}

	if cfg.IsPublishingMode {
		svc.HomePageClient = homepage.NewPublishingClient(svc.Clients, cfg.SupportedLanguages)
	} else {
		svc.HomePageClient, err = homepage.NewWebClient(ctx, svc.Clients, cfg.CacheUpdateInterval, cfg.SupportedLanguages)
		if err != nil {
			log.Fatal(ctx, "failed to create homepage web client", err)
			return err
		}
	}
	if cfg.EnableNewNavBar {
		if err := svc.HomePageClient.AddNavigationCache(ctx, cfg.CacheNavigationUpdateInterval); err != nil {
			log.Fatal(ctx, "failed to add navigation cache to homepage client", err)
			return err
		}
	}

	// Start background polling of topics API for navbar data (changes)
	go svc.HomePageClient.StartBackgroundUpdate(ctx, svcErrors)

	if cfg.EnableCensusTopicSubsection {
		// Initialise caching census topics
		cache.CensusTopicID = cfg.CensusTopicID
		svc.Cache.CensusTopic, err = cache.NewTopicCache(ctx, &cfg.CacheCensusTopicUpdateInterval)
		if err != nil {
			log.Error(ctx, "failed to create topics cache", err)
			return err
		}

		if cfg.IsPublishingMode {
			if err = svc.Cache.CensusTopic.AddUpdateFunc(ctx, cache.CensusTopicID, cachePrivate.UpdateCensusTopic(ctx, cfg.CensusTopicID, cfg.ServiceAuthToken, svc.Clients.Topic)); err != nil {
				log.Error(ctx, "failed to create topics cache", err)
				return err
			}
		} else {
			if err = svc.Cache.CensusTopic.AddUpdateFunc(ctx, cache.CensusTopicID, cachePublic.UpdateCensusTopic(ctx, cfg.CensusTopicID, svc.Clients.Topic)); err != nil {
				log.Error(ctx, "failed to create topics cache", err)
				return err
			}
		}

		go svc.Cache.CensusTopic.StartUpdates(ctx, svcErrors)
	}

	return nil
}
