package homepage

import (
	"context"
	"github.com/ONSdigital/dp-frontend-homepage-controller/model"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	coreModel "github.com/ONSdigital/dp-renderer/model"
	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ReneKroon/ttlcache"

	"github.com/ONSdigital/dp-api-clients-go/v2/image"
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

type mockClientError struct{}

func (e *mockClientError) Error() string { return "client error" }
func (e *mockClientError) Code() int     { return http.StatusNotFound }

<<<<<<< HEAD
// doTestRequest helper function that creates a router and mocks requests
func doTestRequest(target string, req *http.Request, handlerFunc http.HandlerFunc, w *httptest.ResponseRecorder) *httptest.ResponseRecorder {
	if w == nil {
		w = httptest.NewRecorder()
	}
	router := mux.NewRouter()
	router.Path(target).HandlerFunc(handlerFunc)
	router.ServeHTTP(w, req)
	return w
}
=======
	mockedHomepageData := zebedee.HomepageContent{
		Intro: zebedee.Intro{
			Title:    "Welcome to the Office for National Statistics",
			Markdown: "Markdown text here",
		},
		FeaturedContent: []zebedee.Featured{
			{
				Title:       "Featured content one",
				Description: "Featured content one description",
				URI:         "Featured content one URI",
				ImageID:     "123",
			},
			{
				Title:       "Featured content two",
				Description: "Featured content two description",
				URI:         "Featured content two URI",
				ImageID:     "456",
			},
			{
				Title:       "Featured content three",
				Description: "Featured content three description",
				URI:         "Featured content three URI",
				ImageID:     "",
			},
		},
		ServiceMessage: "",
		URI:            "/",
		Type:           "",
		Description:    zebedee.HomepageDescription{},
	}
	var mockedImageDownloadData = map[string]image.ImageDownload{
		"123": {
			Size:  111111,
			Type:  "blah",
			Href:  "http://www.example.com/images/123/original.png",
			State: "completed",
		},
		"456": {
			Size:  111111,
			Type:  "blah",
			Href:  "http://www.example.com/images/456/original.png",
			State: "completed",
		},
	}
	expectedSuccessResponse := "<html><body><h1>Some HTML from renderer!</h1></body></html>"
>>>>>>> develop

var (
	userAccessToken string
	collectionID string
	lang string
)

<<<<<<< HEAD
func TestUnitHomepageHandlerSuccess(t *testing.T) {
	t.Parallel()

	Convey("Given a valid request", t, func() {
		req := httptest.NewRequest("GET", "/", nil)

		cfg, err := config.Get()
		So(err, ShouldBeNil)

		mockedRendererClient := &RenderClientMock{
			BuildPageFunc: func(w io.Writer, pageModel interface{}, templateName string) {},
			NewBasePageModelFunc: func() coreModel.Page {
				return coreModel.Page{}
=======
		mockImageClient := &ImageClientMock{
			GetDownloadVariantFunc: func(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, imageID, variant string) (image.ImageDownload, error) {
				return mockedImageDownloadData[imageID], nil
			},
			CheckerFunc: func(ctx context.Context, check *health.CheckState) error {
				return nil
>>>>>>> develop
			},
		}

		mockedHomepageClienter := &HomepageClienterMock{
			CloseFunc: func()  {},
			GetHomePageFunc: func(ctx context.Context, userAccessToken string, collectionID string, lang string) (*model.HomepageData, error) {
				return &model.HomepageData{}, nil
			},
			StartBackgroundUpdateFunc: func(ctx context.Context, errorChannel chan error)  {},
		}

<<<<<<< HEAD
		Convey("When Read is called", func() {
			w := doTestRequest("/", req, Handler(cfg, mockedHomepageClienter, mockedRendererClient), nil)

			Convey("Then a 200 OK status should be returned", func() {
				So(w.Code, ShouldEqual, http.StatusOK)

				So(len(mockedRendererClient.BuildPageCalls()), ShouldEqual, 1)

			})
		})
	})
}
=======
		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("X-Florence-Token", "testuser")
		rec := httptest.NewRecorder()
		router := mux.NewRouter()
		cache := ttlcache.NewCache()
		cache.SetTTL(10 * time.Millisecond)
		clients := &Clients{
			Renderer: mockRenderClient,
			Zebedee:  mockZebedeeClient,
			ImageAPI: mockImageClient,
		}
		homepageClient := NewHomePagePublishingClient(clients)
		router.Path("/").HandlerFunc(Handler(homepageClient))

		Convey("returns 200 response", func() {
			router.ServeHTTP(rec, req)
			So(rec.Code, ShouldEqual, http.StatusOK)
		})

		Convey("renderer returns HTML body", func() {
			router.ServeHTTP(rec, req)
			response := rec.Body.String()
			So(response, ShouldEqual, expectedSuccessResponse)
		})
	})
}
>>>>>>> develop
