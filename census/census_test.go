package census

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/dp-frontend-homepage-controller/cache"
	"github.com/ONSdigital/dp-frontend-homepage-controller/config"
	"github.com/ONSdigital/dp-frontend-homepage-controller/homepage"
	"github.com/ONSdigital/dp-frontend-homepage-controller/model"
	coreModel "github.com/ONSdigital/dp-renderer/model"
	topicModel "github.com/ONSdigital/dp-topic-api/models"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

type mockClientError struct{}

func (e *mockClientError) Error() string { return "client error" }
func (e *mockClientError) Code() int     { return http.StatusNotFound }

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

func TestUnitCensusHandlerSuccess(t *testing.T) {
	t.Parallel()

	Convey("Given a valid request", t, func() {
		req := httptest.NewRequest("GET", "/census", nil)

		cfg, err := config.Get()
		So(err, ShouldBeNil)

		mockedRendererClient := &RenderClientMock{
			BuildPageFunc: func(w io.Writer, pageModel interface{}, templateName string) {},
			NewBasePageModelFunc: func() coreModel.Page {
				return coreModel.Page{}
			},
		}

		mockedHomepageClienter := &homepage.ClienterMock{
			CloseFunc: func() {},
			GetHomePageFunc: func(ctx context.Context, userAccessToken string, collectionID string, lang string) (*model.HomepageData, error) {
				return &model.HomepageData{}, nil
			},
			GetNavigationDataFunc: func(ctx context.Context, lang string) (*topicModel.Navigation, error) {
				return &topicModel.Navigation{}, nil
			},
			StartBackgroundUpdateFunc: func(ctx context.Context, errorChannel chan error) {},
		}

		Convey("When Read is called", func() {
			w := doTestRequest("/census", req, Handler(cfg, cache.List{}, mockedHomepageClienter, mockedRendererClient), nil)

			Convey("Then a 200 OK status should be returned", func() {
				So(w.Code, ShouldEqual, http.StatusOK)

				So(len(mockedRendererClient.BuildPageCalls()), ShouldEqual, 1)
			})
		})
	})
}
