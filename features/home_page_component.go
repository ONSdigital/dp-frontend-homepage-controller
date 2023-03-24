package feature

import (
	"context"
	"io"
	"net/http"

	"github.com/ONSdigital/dp-api-clients-go/v2/health"
	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	"github.com/ONSdigital/dp-frontend-homepage-controller/homepage"
	"github.com/ONSdigital/dp-frontend-homepage-controller/service"
	"github.com/ONSdigital/dp-frontend-homepage-controller/service/mock"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	dphttp "github.com/ONSdigital/dp-net/v2/http"
	coreModel "github.com/ONSdigital/dp-renderer/model"
	top "github.com/ONSdigital/dp-topic-api/sdk/mocks"
	"github.com/ONSdigital/log.go/log"
	"github.com/cucumber/godog"
)

type HomePageComponent struct {
	Config         *config.Config
	errorChan      chan error
	HTTPServer     *http.Server
	ServiceRunning bool
	svc            *service.Service
}

// Mocks
type HealthCheckerMock struct{}

func NewHealthCheckerMock() *HealthCheckerMock {
	return &HealthCheckerMock{}
}
func (h *HealthCheckerMock) Start(ctx context.Context) {}
func (h *HealthCheckerMock) Stop()                     {}
func (h *HealthCheckerMock) AddCheck(name string, checker healthcheck.Checker) (err error) {
	return nil
}
func (h *HealthCheckerMock) Handler(w http.ResponseWriter, req *http.Request) {}

func New(ctx context.Context) (*HomePageComponent, error) {
	svcErrors := make(chan error, 1)
	c := &HomePageComponent{
		errorChan:      svcErrors,
		HTTPServer:     &http.Server{},
		ServiceRunning: false,
	}

	var err error

	c.Config, err = config.Get()
	if err != nil {
		return nil, err
	}

	initMock := &mock.InitialiserMock{
		DoGetHTTPServerFunc:   c.DoGetHTTPServer,
		DoGetHealthClientFunc: c.DoGetHealthClient,
		DoGetHealthCheckFunc:  c.DoGetHealthCheck,
	}

	svcList := service.NewServiceList(initMock)
	cfg := c.Config

	c.svc = service.New()
	c.svc.InitiateServiceList(cfg, svcList)
	cfg.SiteDomain = "localhost"

	rendererClientMock := &homepage.RenderClientMock{
		BuildPageFunc: func(w io.Writer, pageModel interface{}, templateName string) {},
		NewBasePageModelFunc: func() coreModel.Page {
			return coreModel.Page{}
		},
	}

	c.svc.Clients = &homepage.Clients{
		Zebedee: &homepage.ZebedeeClientMock{
			CheckerFunc:                 nil,
			GetHomepageContentFunc:      GetHomepageContentFuncMock,
			GetTimeseriesMainFigureFunc: GetTimeseriesMainFigureFuncMock,
		},
		ImageAPI: &homepage.ImageClientMock{
			CheckerFunc:            nil,
			GetDownloadVariantFunc: GetDownloadVariantFuncMock,
		},
		Renderer: rendererClientMock,
		Topic: &top.ClienterMock{
			CheckerFunc:              nil,
			GetNavigationPublicFunc:  nil,
			GetRootTopicsPrivateFunc: nil,
			GetRootTopicsPublicFunc:  nil,
			GetSubtopicsPrivateFunc:  nil,
			GetSubtopicsPublicFunc:   GetSubtopicsPublicFuncMock,
			GetTopicPrivateFunc:      nil,
			GetTopicPublicFunc:       GetTopicPublicFuncMock,
			HealthFunc:               nil,
			URLFunc:                  nil,
		},
	}
	err = c.svc.Init(ctx, cfg, svcList, "1", "", "", svcErrors)
	if err != nil {
		return nil, err
	}
	err = c.svc.Run(ctx, cfg, svcList, svcErrors)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *HomePageComponent) StopService(ctx context.Context) {
	err := c.svc.Close(ctx)
	if err != nil {
		log.Error(err)
	}
}

func (c *HomePageComponent) DoGetHTTPServer(bindAddr string, router http.Handler) service.HTTPServer {
	s := dphttp.NewServer(bindAddr, router)
	s.HandleOSSignals = false
	return s
}

func (c *HomePageComponent) DoGetHealthClient(name, url string) *health.Client {
	return nil
}

func (c *HomePageComponent) DoGetHealthCheck(cfg *config.Config, buildTime, gitCommit, version string) (service.HealthChecker, error) {
	// HealthChecker
	return NewHealthCheckerMock(), nil
}

func (c *HomePageComponent) RegisterSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^the census hub flags are enabled`, c.CensusHubFlagsAreEnabled)
}

func (c *HomePageComponent) CensusHubFlagsAreEnabled() error {
	c.Config.EnableCustomDataset = true
	c.Config.EnableGetDataCard = true
	c.Config.DatasetFinderEnabled = true
	return nil
}
