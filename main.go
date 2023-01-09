package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"

	"github.com/ONSdigital/dp-api-clients-go/v2/image"
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	"github.com/ONSdigital/dp-frontend-homepage-controller/homepage"
	"github.com/ONSdigital/dp-frontend-homepage-controller/service"
	topicCli "github.com/ONSdigital/dp-topic-api/sdk"
	"github.com/ONSdigital/log.go/v2/log"
)

const serviceName = "dp-frontend-homepage-controller"

var (
	// BuildTime represents the time in which the service was built
	BuildTime string
	// GitCommit represents the commit (SHA-1) hash of the service that is running
	GitCommit string
	// Version represents the version of the service that is running
	Version string
)

func main() {
	log.Namespace = serviceName
	ctx := context.Background()

	if err := run(ctx); err != nil {
		log.Fatal(ctx, "unable to run application", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Create service initializer and an error channel for fatal errors
	svcErrors := make(chan error, 1)
	svcList := service.NewServiceList(&service.Init{})

	// Read config
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(ctx, "unable to retrieve service configuration", err)
		return err
	}
	log.Info(ctx, "got service configuration", log.Data{"config": cfg})

	// Start service
	svc := service.New()
	svc.InitiateServiceList(cfg, svcList)
	svc.Clients = &homepage.Clients{
		Zebedee:  zebedee.NewWithHealthClient(svc.RouterHealthClient),
		ImageAPI: image.NewWithHealthClient(svc.RouterHealthClient),
		Topic:    topicCli.NewWithHealthClient(svc.RouterHealthClient),
	}

	if err := svc.Init(ctx, cfg, svcList, BuildTime, GitCommit, Version, svcErrors); err != nil {
		return errors.Wrap(err, "running service failed")
	}

	err = svc.Run(ctx, cfg, svcList, svcErrors)
	if err != nil {
		return errors.Wrap(err, "running service failed")
	}

	// Blocks until an os interrupt or a fatal error occurs
	select {
	case err := <-svcErrors:
		log.Error(ctx, "service error received", err)
	case sig := <-signals:
		log.Info(ctx, "os signal received", log.Data{"signal": sig})
	}
	return svc.Close(ctx)
}
