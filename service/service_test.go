package service_test

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/ONSdigital/dp-api-clients-go/v2/health"
	"github.com/ONSdigital/dp-frontend-homepage-controller/cache"
	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	"github.com/ONSdigital/dp-frontend-homepage-controller/homepage"
	"github.com/ONSdigital/dp-frontend-homepage-controller/service"
	"github.com/ONSdigital/dp-frontend-homepage-controller/service/mock"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	dphttp "github.com/ONSdigital/dp-net/v3/http"
	"github.com/ONSdigital/dp-topic-api/models"
	"github.com/ONSdigital/dp-topic-api/sdk"
	topicErrs "github.com/ONSdigital/dp-topic-api/sdk/errors"
	"github.com/ONSdigital/dp-topic-api/sdk/mocks"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	ctx           = context.Background()
	testBuildTime = "BuildTime"
	testGitCommit = "GitCommit"
	testVersion   = "Version"
	errServer     = errors.New("HTTP Server error")
)

var (
	errHealthcheck = errors.New("healthCheck error")
)

var (
	funcDoGetHealthcheckErr = func(cfg *config.Config, buildTime string, gitCommit string, version string) (service.HealthChecker, error) {
		return nil, errHealthcheck
	}

	funcDoGetCacheList = func(ctx context.Context) (cache.List, error) {
		testCensusTopicCache, err := cache.NewTopicCache(ctx, nil)
		if err != nil {
			return cache.List{}, err
		}

		testNavigationCache, err := cache.NewNavigationCache(ctx, nil)
		if err != nil {
			return cache.List{}, err
		}

		cacheList := cache.List{
			CensusTopic: testCensusTopicCache,
			Navigation:  testNavigationCache,
		}

		return cacheList, nil
	}
)

func TestRun(t *testing.T) {
	Convey("Having a set of mocked dependencies", t, func() {
		cfg, err := config.Get()
		cfg.IsPublishingMode = true
		So(err, ShouldBeNil)

		hcMock := &mock.HealthCheckerMock{
			AddCheckFunc: func(name string, checker healthcheck.Checker) error { return nil },
			StartFunc:    func(ctx context.Context) {},
		}

		mockedTopicClienter := &mocks.ClienterMock{
			GetTopicPrivateFunc: func(ctx context.Context, reqHeaders sdk.Headers, id string) (*models.TopicResponse, topicErrs.Error) {
				return &models.TopicResponse{ID: "1234", Next: &models.Topic{ID: "1234", Title: "Census"}}, nil
			},
			GetSubtopicsPrivateFunc: func(ctx context.Context, reqHeaders sdk.Headers, id string) (*models.PrivateSubtopics, topicErrs.Error) {
				return &models.PrivateSubtopics{PrivateItems: &[]models.TopicResponse{{ID: "5678", Next: &models.Topic{ID: "5678", Title: "Age"}}}}, nil
			},
		}

		serverWg := &sync.WaitGroup{}
		serverMock := &mock.HTTPServerMock{
			ListenAndServeFunc: func() error {
				serverWg.Done()
				return nil
			},
		}

		failingServerMock := &mock.HTTPServerMock{
			ListenAndServeFunc: func() error {
				serverWg.Done()
				return errServer
			},
		}

		funcDoGetHealthcheckOk := func(cfg *config.Config, buildTime string, gitCommit string, version string) (service.HealthChecker, error) {
			return hcMock, nil
		}

		funcDoGetHTTPServer := func(bindAddr string, router http.Handler) service.HTTPServer {
			return serverMock
		}

		funcDoGetFailingHTTPSerer := func(bindAddr string, router http.Handler) service.HTTPServer {
			return failingServerMock
		}

		funcDoGetHealthClientOk := func(name string, url string) *health.Client {
			return &health.Client{
				URL:    url,
				Name:   name,
				Client: newMockHTTPClient(&http.Response{}, nil),
			}
		}

		cacheList, err := funcDoGetCacheList(ctx)
		So(err, ShouldBeNil)

		Convey("Given that initialising Healthcheck returns an error", func() {
			initMock := &mock.InitialiserMock{
				DoGetHealthClientFunc: funcDoGetHealthClientOk,
				DoGetHealthCheckFunc:  funcDoGetHealthcheckErr,
			}
			mockSvcList := service.NewServiceList(initMock)

			svcErrors := make(chan error, 1)
			svc := service.New()
			svc.Clients = &homepage.Clients{
				Topic: mockedTopicClienter,
			}

			svc.InitiateServiceList(cfg, mockSvcList)

			err := svc.Init(ctx, cfg, mockSvcList, testBuildTime, testGitCommit, testVersion, svcErrors)
			Convey("Then service Run fails with the same error and the flag is not set. No further initialisations are attempted", func() {
				So(err, ShouldResemble, errHealthcheck)
				So(mockSvcList.HealthCheck, ShouldBeFalse)
			})
		})

		Convey("Given that Checkers cannot be registered", func() {
			errAddheckFail := errors.New("Error(s) registering checkers for healthcheck")
			hcMockAddFail := &mock.HealthCheckerMock{
				AddCheckFunc: func(name string, checker healthcheck.Checker) error { return errAddheckFail },
				StartFunc:    func(ctx context.Context) {},
			}

			initMock := &mock.InitialiserMock{
				DoGetHealthClientFunc: funcDoGetHealthClientOk,
				DoGetHealthCheckFunc: func(cfg *config.Config, buildTime string, gitCommit string, version string) (service.HealthChecker, error) {
					return hcMockAddFail, nil
				},
			}
			svcErrors := make(chan error, 1)
			svcList := service.NewServiceList(initMock)
			svc := service.New()
			svc.Clients = &homepage.Clients{
				Topic: mockedTopicClienter,
			}

			err := svc.Init(ctx, cfg, svcList, testBuildTime, testGitCommit, testVersion, svcErrors)

			Convey("Then service Run fails, but all checks try to register", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldResemble, fmt.Sprintf("unable to register checkers: %s", errAddheckFail.Error()))
				So(svcList.HealthCheck, ShouldBeTrue)
				So(len(hcMockAddFail.AddCheckCalls()), ShouldEqual, 1)
				So(hcMockAddFail.AddCheckCalls()[0].Name, ShouldResemble, "API router")
			})
		})

		Convey("Given that all dependencies are successfully initialised", func() {
			initMock := &mock.InitialiserMock{
				DoGetHealthClientFunc: funcDoGetHealthClientOk,
				DoGetHealthCheckFunc:  funcDoGetHealthcheckOk,
				DoGetHTTPServerFunc:   funcDoGetHTTPServer,
			}
			svcErrors := make(chan error, 1)
			serverWg.Add(1)
			svc := service.New()
			svc.Cache = cacheList
			mockSvcList := service.NewServiceList(initMock)
			svc.InitiateServiceList(cfg, mockSvcList)
			svc.Clients = &homepage.Clients{
				Topic: mockedTopicClienter,
			}

			err := svc.Init(ctx, cfg, mockSvcList, testBuildTime, testGitCommit, testVersion, svcErrors)
			err = svc.Run(ctx, cfg, mockSvcList, svcErrors)
			Convey("Then service Run succeeds and all the flags are set", func() {
				So(err, ShouldBeNil)
				So(mockSvcList.HealthCheck, ShouldBeTrue)
			})

			Convey("The checkers are registered and the healthcheck and http server started", func() {
				So(len(hcMock.AddCheckCalls()), ShouldEqual, 1)
				So(hcMock.AddCheckCalls()[0].Name, ShouldResemble, "API router")
				So(len(initMock.DoGetHTTPServerCalls()), ShouldEqual, 1)
				So(initMock.DoGetHTTPServerCalls()[0].BindAddr, ShouldEqual, ":24400")
				So(len(hcMock.StartCalls()), ShouldEqual, 1)
				serverWg.Wait() // Wait for HTTP server go-routine to finish
				So(len(serverMock.ListenAndServeCalls()), ShouldEqual, 1)
			})
		})
		Convey("Given that all dependencies are successfully initialised but the http server fails", func() {
			initMock := &mock.InitialiserMock{
				DoGetHealthClientFunc: funcDoGetHealthClientOk,
				DoGetHealthCheckFunc:  funcDoGetHealthcheckOk,
				DoGetHTTPServerFunc:   funcDoGetFailingHTTPSerer,
			}
			svcErrors := make(chan error, 1)
			serverWg.Add(1)
			svc := service.New()

			mockSvcList := service.NewServiceList(initMock)
			svc.InitiateServiceList(cfg, mockSvcList)
			svc.Clients = &homepage.Clients{
				Topic: mockedTopicClienter,
			}

			err = svc.Init(ctx, cfg, mockSvcList, testBuildTime, testGitCommit, testVersion, svcErrors)
			So(err, ShouldBeNil)
			err = svc.Run(ctx, cfg, mockSvcList, svcErrors)

			Convey("Then the ersror is returned in the error channel", func() {
				sErr := <-svcErrors
				So(sErr.Error(), ShouldResemble, fmt.Sprintf("failure in http listen and serve: %s", errServer.Error()))
				So(len(failingServerMock.ListenAndServeCalls()), ShouldEqual, 1)
			})
		})
	})
}

func TestClose(t *testing.T) {
	Convey("Having a correctly initialised service", t, func() {
		cfg, err := config.Get()
		So(err, ShouldBeNil)

		hcStopped := false

		// healthcheck Stop does not depend on any other service being closed/stopped
		hcMock := &mock.HealthCheckerMock{
			AddCheckFunc: func(name string, checker healthcheck.Checker) error { return nil },
			StartFunc:    func(ctx context.Context) {},
			StopFunc:     func() { hcStopped = true },
		}

		// server Shutdown will fail if healthcheck is not stopped
		serverMock := &mock.HTTPServerMock{
			ListenAndServeFunc: func() error { return nil },
			ShutdownFunc: func(ctx context.Context) error {
				if !hcStopped {
					return errors.New("Server stopped before healthcheck")
				}
				return nil
			},
		}

		Convey("Closing the service results in all the dependencies being closed in the expected order", func() {
			svcList := service.NewServiceList(nil)
			svcList.HealthCheck = true
			svc := service.Service{
				Config:      cfg,
				ServiceList: svcList,
				Server:      serverMock,
				HealthCheck: hcMock,
			}
			err = svc.Close(context.Background())
			So(err, ShouldBeNil)
			So(len(hcMock.StopCalls()), ShouldEqual, 1)
			So(len(serverMock.ShutdownCalls()), ShouldEqual, 1)
		})

		Convey("If services fail to stop, the Close operation tries to close all dependencies and returns an error", func() {
			failingserverMock := &mock.HTTPServerMock{
				ListenAndServeFunc: func() error { return nil },
				ShutdownFunc: func(ctx context.Context) error {
					return errors.New("Failed to stop http server")
				},
			}

			svcList := service.NewServiceList(nil)
			svcList.HealthCheck = true
			svc := service.Service{
				Config:      cfg,
				ServiceList: svcList,
				Server:      failingserverMock,
				HealthCheck: hcMock,
			}
			err = svc.Close(context.Background())
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, "failed to shutdown gracefully")
			So(len(hcMock.StopCalls()), ShouldEqual, 1)
			So(len(failingserverMock.ShutdownCalls()), ShouldEqual, 1)
		})

		Convey("If service times out while shutting down, the Close operation fails with the expected error", func() {
			cfg.GracefulShutdownTimeout = 1 * time.Millisecond
			timeoutServerMock := &mock.HTTPServerMock{
				ListenAndServeFunc: func() error { return nil },
				ShutdownFunc: func(ctx context.Context) error {
					time.Sleep(100 * time.Millisecond)
					return nil
				},
			}

			svcList := service.NewServiceList(nil)
			svcList.HealthCheck = true
			svc := service.Service{
				Config:      cfg,
				ServiceList: svcList,
				Server:      timeoutServerMock,
				HealthCheck: hcMock,
			}
			err = svc.Close(context.Background())
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, "context deadline exceeded")
			So(len(hcMock.StopCalls()), ShouldEqual, 1)
			So(len(timeoutServerMock.ShutdownCalls()), ShouldEqual, 1)
		})
	})
}

func newMockHTTPClient(r *http.Response, err error) *dphttp.ClienterMock {
	return &dphttp.ClienterMock{
		SetPathsWithNoRetriesFunc: func(paths []string) {},
		GetPathsWithNoRetriesFunc: func() []string { return []string{} },
		DoFunc: func(ctx context.Context, req *http.Request) (*http.Response, error) {
			return r, err
		},
	}
}
