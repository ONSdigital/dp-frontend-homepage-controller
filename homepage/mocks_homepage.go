// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package homepage

import (
	"context"
	"github.com/ONSdigital/dp-api-clients-go/v2/image"
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/dp-renderer/model"
	"io"
	"sync"
)

// Ensure, that ZebedeeClientMock does implement ZebedeeClient.
// If this is not the case, regenerate this file with moq.
var _ ZebedeeClient = &ZebedeeClientMock{}

// ZebedeeClientMock is a mock implementation of ZebedeeClient.
//
// 	func TestSomethingThatUsesZebedeeClient(t *testing.T) {
//
// 		// make and configure a mocked ZebedeeClient
// 		mockedZebedeeClient := &ZebedeeClientMock{
// 			CheckerFunc: func(ctx context.Context, check *health.CheckState) error {
// 				panic("mock out the Checker method")
// 			},
// 			GetHomepageContentFunc: func(ctx context.Context, userAccessToken string, collectionID string, lang string, path string) (zebedee.HomepageContent, error) {
// 				panic("mock out the GetHomepageContent method")
// 			},
// 			GetTimeseriesMainFigureFunc: func(ctx context.Context, userAuthToken string, collectionID string, lang string, uri string) (zebedee.TimeseriesMainFigure, error) {
// 				panic("mock out the GetTimeseriesMainFigure method")
// 			},
// 		}
//
// 		// use mockedZebedeeClient in code that requires ZebedeeClient
// 		// and then make assertions.
//
// 	}
type ZebedeeClientMock struct {
	// CheckerFunc mocks the Checker method.
	CheckerFunc func(ctx context.Context, check *health.CheckState) error

	// GetHomepageContentFunc mocks the GetHomepageContent method.
	GetHomepageContentFunc func(ctx context.Context, userAccessToken string, collectionID string, lang string, path string) (zebedee.HomepageContent, error)

	// GetTimeseriesMainFigureFunc mocks the GetTimeseriesMainFigure method.
	GetTimeseriesMainFigureFunc func(ctx context.Context, userAuthToken string, collectionID string, lang string, uri string) (zebedee.TimeseriesMainFigure, error)

	// calls tracks calls to the methods.
	calls struct {
		// Checker holds details about calls to the Checker method.
		Checker []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Check is the check argument value.
			Check *health.CheckState
		}
		// GetHomepageContent holds details about calls to the GetHomepageContent method.
		GetHomepageContent []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserAccessToken is the userAccessToken argument value.
			UserAccessToken string
			// CollectionID is the collectionID argument value.
			CollectionID string
			// Lang is the lang argument value.
			Lang string
			// Path is the path argument value.
			Path string
		}
		// GetTimeseriesMainFigure holds details about calls to the GetTimeseriesMainFigure method.
		GetTimeseriesMainFigure []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserAuthToken is the userAuthToken argument value.
			UserAuthToken string
			// CollectionID is the collectionID argument value.
			CollectionID string
			// Lang is the lang argument value.
			Lang string
			// URI is the uri argument value.
			URI string
		}
	}
	lockChecker                 sync.RWMutex
	lockGetHomepageContent      sync.RWMutex
	lockGetTimeseriesMainFigure sync.RWMutex
}

// Checker calls CheckerFunc.
func (mock *ZebedeeClientMock) Checker(ctx context.Context, check *health.CheckState) error {
	if mock.CheckerFunc == nil {
		panic("ZebedeeClientMock.CheckerFunc: method is nil but ZebedeeClient.Checker was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Check *health.CheckState
	}{
		Ctx:   ctx,
		Check: check,
	}
	mock.lockChecker.Lock()
	mock.calls.Checker = append(mock.calls.Checker, callInfo)
	mock.lockChecker.Unlock()
	return mock.CheckerFunc(ctx, check)
}

// CheckerCalls gets all the calls that were made to Checker.
// Check the length with:
//     len(mockedZebedeeClient.CheckerCalls())
func (mock *ZebedeeClientMock) CheckerCalls() []struct {
	Ctx   context.Context
	Check *health.CheckState
} {
	var calls []struct {
		Ctx   context.Context
		Check *health.CheckState
	}
	mock.lockChecker.RLock()
	calls = mock.calls.Checker
	mock.lockChecker.RUnlock()
	return calls
}

// GetHomepageContent calls GetHomepageContentFunc.
func (mock *ZebedeeClientMock) GetHomepageContent(ctx context.Context, userAccessToken string, collectionID string, lang string, path string) (zebedee.HomepageContent, error) {
	if mock.GetHomepageContentFunc == nil {
		panic("ZebedeeClientMock.GetHomepageContentFunc: method is nil but ZebedeeClient.GetHomepageContent was just called")
	}
	callInfo := struct {
		Ctx             context.Context
		UserAccessToken string
		CollectionID    string
		Lang            string
		Path            string
	}{
		Ctx:             ctx,
		UserAccessToken: userAccessToken,
		CollectionID:    collectionID,
		Lang:            lang,
		Path:            path,
	}
	mock.lockGetHomepageContent.Lock()
	mock.calls.GetHomepageContent = append(mock.calls.GetHomepageContent, callInfo)
	mock.lockGetHomepageContent.Unlock()
	return mock.GetHomepageContentFunc(ctx, userAccessToken, collectionID, lang, path)
}

// GetHomepageContentCalls gets all the calls that were made to GetHomepageContent.
// Check the length with:
//     len(mockedZebedeeClient.GetHomepageContentCalls())
func (mock *ZebedeeClientMock) GetHomepageContentCalls() []struct {
	Ctx             context.Context
	UserAccessToken string
	CollectionID    string
	Lang            string
	Path            string
} {
	var calls []struct {
		Ctx             context.Context
		UserAccessToken string
		CollectionID    string
		Lang            string
		Path            string
	}
	mock.lockGetHomepageContent.RLock()
	calls = mock.calls.GetHomepageContent
	mock.lockGetHomepageContent.RUnlock()
	return calls
}

// GetTimeseriesMainFigure calls GetTimeseriesMainFigureFunc.
func (mock *ZebedeeClientMock) GetTimeseriesMainFigure(ctx context.Context, userAuthToken string, collectionID string, lang string, uri string) (zebedee.TimeseriesMainFigure, error) {
	if mock.GetTimeseriesMainFigureFunc == nil {
		panic("ZebedeeClientMock.GetTimeseriesMainFigureFunc: method is nil but ZebedeeClient.GetTimeseriesMainFigure was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		UserAuthToken string
		CollectionID  string
		Lang          string
		URI           string
	}{
		Ctx:           ctx,
		UserAuthToken: userAuthToken,
		CollectionID:  collectionID,
		Lang:          lang,
		URI:           uri,
	}
	mock.lockGetTimeseriesMainFigure.Lock()
	mock.calls.GetTimeseriesMainFigure = append(mock.calls.GetTimeseriesMainFigure, callInfo)
	mock.lockGetTimeseriesMainFigure.Unlock()
	return mock.GetTimeseriesMainFigureFunc(ctx, userAuthToken, collectionID, lang, uri)
}

// GetTimeseriesMainFigureCalls gets all the calls that were made to GetTimeseriesMainFigure.
// Check the length with:
//     len(mockedZebedeeClient.GetTimeseriesMainFigureCalls())
func (mock *ZebedeeClientMock) GetTimeseriesMainFigureCalls() []struct {
	Ctx           context.Context
	UserAuthToken string
	CollectionID  string
	Lang          string
	URI           string
} {
	var calls []struct {
		Ctx           context.Context
		UserAuthToken string
		CollectionID  string
		Lang          string
		URI           string
	}
	mock.lockGetTimeseriesMainFigure.RLock()
	calls = mock.calls.GetTimeseriesMainFigure
	mock.lockGetTimeseriesMainFigure.RUnlock()
	return calls
}

// Ensure, that ImageClientMock does implement ImageClient.
// If this is not the case, regenerate this file with moq.
var _ ImageClient = &ImageClientMock{}

// ImageClientMock is a mock implementation of ImageClient.
//
// 	func TestSomethingThatUsesImageClient(t *testing.T) {
//
// 		// make and configure a mocked ImageClient
// 		mockedImageClient := &ImageClientMock{
// 			CheckerFunc: func(ctx context.Context, check *health.CheckState) error {
// 				panic("mock out the Checker method")
// 			},
// 			GetDownloadVariantFunc: func(ctx context.Context, userAuthToken string, serviceAuthToken string, collectionID string, imageID string, variant string) (image.ImageDownload, error) {
// 				panic("mock out the GetDownloadVariant method")
// 			},
// 		}
//
// 		// use mockedImageClient in code that requires ImageClient
// 		// and then make assertions.
//
// 	}
type ImageClientMock struct {
	// CheckerFunc mocks the Checker method.
	CheckerFunc func(ctx context.Context, check *health.CheckState) error

	// GetDownloadVariantFunc mocks the GetDownloadVariant method.
	GetDownloadVariantFunc func(ctx context.Context, userAuthToken string, serviceAuthToken string, collectionID string, imageID string, variant string) (image.ImageDownload, error)

	// calls tracks calls to the methods.
	calls struct {
		// Checker holds details about calls to the Checker method.
		Checker []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Check is the check argument value.
			Check *health.CheckState
		}
		// GetDownloadVariant holds details about calls to the GetDownloadVariant method.
		GetDownloadVariant []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserAuthToken is the userAuthToken argument value.
			UserAuthToken string
			// ServiceAuthToken is the serviceAuthToken argument value.
			ServiceAuthToken string
			// CollectionID is the collectionID argument value.
			CollectionID string
			// ImageID is the imageID argument value.
			ImageID string
			// Variant is the variant argument value.
			Variant string
		}
	}
	lockChecker            sync.RWMutex
	lockGetDownloadVariant sync.RWMutex
}

// Checker calls CheckerFunc.
func (mock *ImageClientMock) Checker(ctx context.Context, check *health.CheckState) error {
	if mock.CheckerFunc == nil {
		panic("ImageClientMock.CheckerFunc: method is nil but ImageClient.Checker was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Check *health.CheckState
	}{
		Ctx:   ctx,
		Check: check,
	}
	mock.lockChecker.Lock()
	mock.calls.Checker = append(mock.calls.Checker, callInfo)
	mock.lockChecker.Unlock()
	return mock.CheckerFunc(ctx, check)
}

// CheckerCalls gets all the calls that were made to Checker.
// Check the length with:
//     len(mockedImageClient.CheckerCalls())
func (mock *ImageClientMock) CheckerCalls() []struct {
	Ctx   context.Context
	Check *health.CheckState
} {
	var calls []struct {
		Ctx   context.Context
		Check *health.CheckState
	}
	mock.lockChecker.RLock()
	calls = mock.calls.Checker
	mock.lockChecker.RUnlock()
	return calls
}

// GetDownloadVariant calls GetDownloadVariantFunc.
func (mock *ImageClientMock) GetDownloadVariant(ctx context.Context, userAuthToken string, serviceAuthToken string, collectionID string, imageID string, variant string) (image.ImageDownload, error) {
	if mock.GetDownloadVariantFunc == nil {
		panic("ImageClientMock.GetDownloadVariantFunc: method is nil but ImageClient.GetDownloadVariant was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		UserAuthToken    string
		ServiceAuthToken string
		CollectionID     string
		ImageID          string
		Variant          string
	}{
		Ctx:              ctx,
		UserAuthToken:    userAuthToken,
		ServiceAuthToken: serviceAuthToken,
		CollectionID:     collectionID,
		ImageID:          imageID,
		Variant:          variant,
	}
	mock.lockGetDownloadVariant.Lock()
	mock.calls.GetDownloadVariant = append(mock.calls.GetDownloadVariant, callInfo)
	mock.lockGetDownloadVariant.Unlock()
	return mock.GetDownloadVariantFunc(ctx, userAuthToken, serviceAuthToken, collectionID, imageID, variant)
}

// GetDownloadVariantCalls gets all the calls that were made to GetDownloadVariant.
// Check the length with:
//     len(mockedImageClient.GetDownloadVariantCalls())
func (mock *ImageClientMock) GetDownloadVariantCalls() []struct {
	Ctx              context.Context
	UserAuthToken    string
	ServiceAuthToken string
	CollectionID     string
	ImageID          string
	Variant          string
} {
	var calls []struct {
		Ctx              context.Context
		UserAuthToken    string
		ServiceAuthToken string
		CollectionID     string
		ImageID          string
		Variant          string
	}
	mock.lockGetDownloadVariant.RLock()
	calls = mock.calls.GetDownloadVariant
	mock.lockGetDownloadVariant.RUnlock()
	return calls
}

// Ensure, that RenderClientMock does implement RenderClient.
// If this is not the case, regenerate this file with moq.
var _ RenderClient = &RenderClientMock{}

// RenderClientMock is a mock implementation of RenderClient.
//
// 	func TestSomethingThatUsesRenderClient(t *testing.T) {
//
// 		// make and configure a mocked RenderClient
// 		mockedRenderClient := &RenderClientMock{
// 			BuildPageFunc: func(w io.Writer, pageModel interface{}, templateName string)  {
// 				panic("mock out the BuildPage method")
// 			},
// 			NewBasePageModelFunc: func() model.Page {
// 				panic("mock out the NewBasePageModel method")
// 			},
// 		}
//
// 		// use mockedRenderClient in code that requires RenderClient
// 		// and then make assertions.
//
// 	}
type RenderClientMock struct {
	// BuildPageFunc mocks the BuildPage method.
	BuildPageFunc func(w io.Writer, pageModel interface{}, templateName string)

	// NewBasePageModelFunc mocks the NewBasePageModel method.
	NewBasePageModelFunc func() model.Page

	// calls tracks calls to the methods.
	calls struct {
		// BuildPage holds details about calls to the BuildPage method.
		BuildPage []struct {
			// W is the w argument value.
			W io.Writer
			// PageModel is the pageModel argument value.
			PageModel interface{}
			// TemplateName is the templateName argument value.
			TemplateName string
		}
		// NewBasePageModel holds details about calls to the NewBasePageModel method.
		NewBasePageModel []struct {
		}
	}
	lockBuildPage        sync.RWMutex
	lockNewBasePageModel sync.RWMutex
}

// BuildPage calls BuildPageFunc.
func (mock *RenderClientMock) BuildPage(w io.Writer, pageModel interface{}, templateName string) {
	if mock.BuildPageFunc == nil {
		panic("RenderClientMock.BuildPageFunc: method is nil but RenderClient.BuildPage was just called")
	}
	callInfo := struct {
		W            io.Writer
		PageModel    interface{}
		TemplateName string
	}{
		W:            w,
		PageModel:    pageModel,
		TemplateName: templateName,
	}
	mock.lockBuildPage.Lock()
	mock.calls.BuildPage = append(mock.calls.BuildPage, callInfo)
	mock.lockBuildPage.Unlock()
	mock.BuildPageFunc(w, pageModel, templateName)
}

// BuildPageCalls gets all the calls that were made to BuildPage.
// Check the length with:
//     len(mockedRenderClient.BuildPageCalls())
func (mock *RenderClientMock) BuildPageCalls() []struct {
	W            io.Writer
	PageModel    interface{}
	TemplateName string
} {
	var calls []struct {
		W            io.Writer
		PageModel    interface{}
		TemplateName string
	}
	mock.lockBuildPage.RLock()
	calls = mock.calls.BuildPage
	mock.lockBuildPage.RUnlock()
	return calls
}

// NewBasePageModel calls NewBasePageModelFunc.
func (mock *RenderClientMock) NewBasePageModel() model.Page {
	if mock.NewBasePageModelFunc == nil {
		panic("RenderClientMock.NewBasePageModelFunc: method is nil but RenderClient.NewBasePageModel was just called")
	}
	callInfo := struct {
	}{}
	mock.lockNewBasePageModel.Lock()
	mock.calls.NewBasePageModel = append(mock.calls.NewBasePageModel, callInfo)
	mock.lockNewBasePageModel.Unlock()
	return mock.NewBasePageModelFunc()
}

// NewBasePageModelCalls gets all the calls that were made to NewBasePageModel.
// Check the length with:
//     len(mockedRenderClient.NewBasePageModelCalls())
func (mock *RenderClientMock) NewBasePageModelCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockNewBasePageModel.RLock()
	calls = mock.calls.NewBasePageModel
	mock.lockNewBasePageModel.RUnlock()
	return calls
}
