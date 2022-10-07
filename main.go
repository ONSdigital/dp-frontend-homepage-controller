package main

import (
	"context"
	"github.com/ONSdigital/dp-frontend-homepage-controller/cache"
	"os"
	"os/signal"
	"syscall"

	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	"github.com/ONSdigital/dp-frontend-homepage-controller/service"
	topicCli "github.com/ONSdigital/dp-topic-api/sdk"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/pkg/errors"
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

	// Create service initialiser and an error channel for fatal errors
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
	svc, err := service.Run(ctx, cfg, svcList, BuildTime, GitCommit, Version, svcErrors)
	if err != nil {
		return errors.Wrap(err, "running service failed")
	}

	// Topics caching
	svc.Cache.TopicCache, err = cache.NewTopicCache(ctx, &cfg.CacheCensusTopicUpdateInterval)
	if err != nil {
		log.Error(ctx, "failed to create topics cache", err)
		return err
	}

	topicClient := topicCli.New("http://localhost:25300/topics")

	svc.Cache.TopicCache.AddUpdateFunc(cache.CensusTopicID, cache.UpdateTopic(ctx, topicClient))

	// Start caching
	go svc.Cache.TopicCache.StartUpdates(ctx, svcErrors)

	// Blocks until an os interrupt or a fatal error occurs
	select {
	case err := <-svcErrors:
		log.Error(ctx, "service error received", err)
	case sig := <-signals:
		log.Info(ctx, "os signal received", log.Data{"signal": sig})
	}
	return svc.Close(ctx)
}
