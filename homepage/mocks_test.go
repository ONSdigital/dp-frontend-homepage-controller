// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package homepage

import (
	"context"
	"github.com/ONSdigital/dp-api-clients-go/image"
	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/clients/release_calendar"
	"sync"
)

var (
	lockZebedeeClientMockGetHomepageContent      sync.RWMutex
	lockZebedeeClientMockGetTimeseriesMainFigure sync.RWMutex
)

// Ensure, that ZebedeeClientMock does implement ZebedeeClient.
// If this is not the case, regenerate this file with moq.
var _ ZebedeeClient = &ZebedeeClientMock{}

// ZebedeeClientMock is a mock implementation of ZebedeeClient.
//
//     func TestSomethingThatUsesZebedeeClient(t *testing.T) {
//
//         // make and configure a mocked ZebedeeClient
//         mockedZebedeeClient := &ZebedeeClientMock{
//             GetHomepageContentFunc: func(ctx context.Context, userAccessToken string, path string) (zebedee.HomepageContent, error) {
// 	               panic("mock out the GetHomepageContent method")
//             },
//             GetTimeseriesMainFigureFunc: func(ctx context.Context, userAuthToken string, uri string) (zebedee.TimeseriesMainFigure, error) {
// 	               panic("mock out the GetTimeseriesMainFigure method")
//             },
//         }
//
//         // use mockedZebedeeClient in code that requires ZebedeeClient
//         // and then make assertions.
//
//     }
type ZebedeeClientMock struct {
	// GetHomepageContentFunc mocks the GetHomepageContent method.
	GetHomepageContentFunc func(ctx context.Context, userAccessToken string, path string) (zebedee.HomepageContent, error)

	// GetTimeseriesMainFigureFunc mocks the GetTimeseriesMainFigure method.
	GetTimeseriesMainFigureFunc func(ctx context.Context, userAuthToken string, uri string) (zebedee.TimeseriesMainFigure, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetHomepageContent holds details about calls to the GetHomepageContent method.
		GetHomepageContent []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserAccessToken is the userAccessToken argument value.
			UserAccessToken string
			// Path is the path argument value.
			Path string
		}
		// GetTimeseriesMainFigure holds details about calls to the GetTimeseriesMainFigure method.
		GetTimeseriesMainFigure []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserAuthToken is the userAuthToken argument value.
			UserAuthToken string
			// URI is the uri argument value.
			URI string
		}
	}
}

// GetHomepageContent calls GetHomepageContentFunc.
func (mock *ZebedeeClientMock) GetHomepageContent(ctx context.Context, userAccessToken string, path string) (zebedee.HomepageContent, error) {
	if mock.GetHomepageContentFunc == nil {
		panic("ZebedeeClientMock.GetHomepageContentFunc: method is nil but ZebedeeClient.GetHomepageContent was just called")
	}
	callInfo := struct {
		Ctx             context.Context
		UserAccessToken string
		Path            string
	}{
		Ctx:             ctx,
		UserAccessToken: userAccessToken,
		Path:            path,
	}
	lockZebedeeClientMockGetHomepageContent.Lock()
	mock.calls.GetHomepageContent = append(mock.calls.GetHomepageContent, callInfo)
	lockZebedeeClientMockGetHomepageContent.Unlock()
	return mock.GetHomepageContentFunc(ctx, userAccessToken, path)
}

// GetHomepageContentCalls gets all the calls that were made to GetHomepageContent.
// Check the length with:
//     len(mockedZebedeeClient.GetHomepageContentCalls())
func (mock *ZebedeeClientMock) GetHomepageContentCalls() []struct {
	Ctx             context.Context
	UserAccessToken string
	Path            string
} {
	var calls []struct {
		Ctx             context.Context
		UserAccessToken string
		Path            string
	}
	lockZebedeeClientMockGetHomepageContent.RLock()
	calls = mock.calls.GetHomepageContent
	lockZebedeeClientMockGetHomepageContent.RUnlock()
	return calls
}

// GetTimeseriesMainFigure calls GetTimeseriesMainFigureFunc.
func (mock *ZebedeeClientMock) GetTimeseriesMainFigure(ctx context.Context, userAuthToken string, uri string) (zebedee.TimeseriesMainFigure, error) {
	if mock.GetTimeseriesMainFigureFunc == nil {
		panic("ZebedeeClientMock.GetTimeseriesMainFigureFunc: method is nil but ZebedeeClient.GetTimeseriesMainFigure was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		UserAuthToken string
		URI           string
	}{
		Ctx:           ctx,
		UserAuthToken: userAuthToken,
		URI:           uri,
	}
	lockZebedeeClientMockGetTimeseriesMainFigure.Lock()
	mock.calls.GetTimeseriesMainFigure = append(mock.calls.GetTimeseriesMainFigure, callInfo)
	lockZebedeeClientMockGetTimeseriesMainFigure.Unlock()
	return mock.GetTimeseriesMainFigureFunc(ctx, userAuthToken, uri)
}

// GetTimeseriesMainFigureCalls gets all the calls that were made to GetTimeseriesMainFigure.
// Check the length with:
//     len(mockedZebedeeClient.GetTimeseriesMainFigureCalls())
func (mock *ZebedeeClientMock) GetTimeseriesMainFigureCalls() []struct {
	Ctx           context.Context
	UserAuthToken string
	URI           string
} {
	var calls []struct {
		Ctx           context.Context
		UserAuthToken string
		URI           string
	}
	lockZebedeeClientMockGetTimeseriesMainFigure.RLock()
	calls = mock.calls.GetTimeseriesMainFigure
	lockZebedeeClientMockGetTimeseriesMainFigure.RUnlock()
	return calls
}

var (
	lockBabbageClientMockGetReleaseCalendar sync.RWMutex
)

// Ensure, that BabbageClientMock does implement BabbageClient.
// If this is not the case, regenerate this file with moq.
var _ BabbageClient = &BabbageClientMock{}

// BabbageClientMock is a mock implementation of BabbageClient.
//
//     func TestSomethingThatUsesBabbageClient(t *testing.T) {
//
//         // make and configure a mocked BabbageClient
//         mockedBabbageClient := &BabbageClientMock{
//             GetReleaseCalendarFunc: func(ctx context.Context, userAccessToken string, fromDay string, fromMonth string, fromYear string) (release_calendar.ReleaseCalendar, error) {
// 	               panic("mock out the GetReleaseCalendar method")
//             },
//         }
//
//         // use mockedBabbageClient in code that requires BabbageClient
//         // and then make assertions.
//
//     }
type BabbageClientMock struct {
	// GetReleaseCalendarFunc mocks the GetReleaseCalendar method.
	GetReleaseCalendarFunc func(ctx context.Context, userAccessToken string, fromDay string, fromMonth string, fromYear string) (release_calendar.ReleaseCalendar, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetReleaseCalendar holds details about calls to the GetReleaseCalendar method.
		GetReleaseCalendar []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserAccessToken is the userAccessToken argument value.
			UserAccessToken string
			// FromDay is the fromDay argument value.
			FromDay string
			// FromMonth is the fromMonth argument value.
			FromMonth string
			// FromYear is the fromYear argument value.
			FromYear string
		}
	}
}

// GetReleaseCalendar calls GetReleaseCalendarFunc.
func (mock *BabbageClientMock) GetReleaseCalendar(ctx context.Context, userAccessToken string, fromDay string, fromMonth string, fromYear string) (release_calendar.ReleaseCalendar, error) {
	if mock.GetReleaseCalendarFunc == nil {
		panic("BabbageClientMock.GetReleaseCalendarFunc: method is nil but BabbageClient.GetReleaseCalendar was just called")
	}
	callInfo := struct {
		Ctx             context.Context
		UserAccessToken string
		FromDay         string
		FromMonth       string
		FromYear        string
	}{
		Ctx:             ctx,
		UserAccessToken: userAccessToken,
		FromDay:         fromDay,
		FromMonth:       fromMonth,
		FromYear:        fromYear,
	}
	lockBabbageClientMockGetReleaseCalendar.Lock()
	mock.calls.GetReleaseCalendar = append(mock.calls.GetReleaseCalendar, callInfo)
	lockBabbageClientMockGetReleaseCalendar.Unlock()
	return mock.GetReleaseCalendarFunc(ctx, userAccessToken, fromDay, fromMonth, fromYear)
}

// GetReleaseCalendarCalls gets all the calls that were made to GetReleaseCalendar.
// Check the length with:
//     len(mockedBabbageClient.GetReleaseCalendarCalls())
func (mock *BabbageClientMock) GetReleaseCalendarCalls() []struct {
	Ctx             context.Context
	UserAccessToken string
	FromDay         string
	FromMonth       string
	FromYear        string
} {
	var calls []struct {
		Ctx             context.Context
		UserAccessToken string
		FromDay         string
		FromMonth       string
		FromYear        string
	}
	lockBabbageClientMockGetReleaseCalendar.RLock()
	calls = mock.calls.GetReleaseCalendar
	lockBabbageClientMockGetReleaseCalendar.RUnlock()
	return calls
}

var (
	lockImageClientMockGetImage sync.RWMutex
)

// Ensure, that ImageClientMock does implement ImageClient.
// If this is not the case, regenerate this file with moq.
var _ ImageClient = &ImageClientMock{}

// ImageClientMock is a mock implementation of ImageClient.
//
//     func TestSomethingThatUsesImageClient(t *testing.T) {
//
//         // make and configure a mocked ImageClient
//         mockedImageClient := &ImageClientMock{
//             GetImageFunc: func(ctx context.Context, userAuthToken string, serviceAuthToken string, collectionID string, imageID string) (image.Image, error) {
// 	               panic("mock out the GetImage method")
//             },
//         }
//
//         // use mockedImageClient in code that requires ImageClient
//         // and then make assertions.
//
//     }
type ImageClientMock struct {
	// GetImageFunc mocks the GetImage method.
	GetImageFunc func(ctx context.Context, userAuthToken string, serviceAuthToken string, collectionID string, imageID string) (image.Image, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetImage holds details about calls to the GetImage method.
		GetImage []struct {
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
		}
	}
}

// GetImage calls GetImageFunc.
func (mock *ImageClientMock) GetImage(ctx context.Context, userAuthToken string, serviceAuthToken string, collectionID string, imageID string) (image.Image, error) {
	if mock.GetImageFunc == nil {
		panic("ImageClientMock.GetImageFunc: method is nil but ImageClient.GetImage was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		UserAuthToken    string
		ServiceAuthToken string
		CollectionID     string
		ImageID          string
	}{
		Ctx:              ctx,
		UserAuthToken:    userAuthToken,
		ServiceAuthToken: serviceAuthToken,
		CollectionID:     collectionID,
		ImageID:          imageID,
	}
	lockImageClientMockGetImage.Lock()
	mock.calls.GetImage = append(mock.calls.GetImage, callInfo)
	lockImageClientMockGetImage.Unlock()
	return mock.GetImageFunc(ctx, userAuthToken, serviceAuthToken, collectionID, imageID)
}

// GetImageCalls gets all the calls that were made to GetImage.
// Check the length with:
//     len(mockedImageClient.GetImageCalls())
func (mock *ImageClientMock) GetImageCalls() []struct {
	Ctx              context.Context
	UserAuthToken    string
	ServiceAuthToken string
	CollectionID     string
	ImageID          string
} {
	var calls []struct {
		Ctx              context.Context
		UserAuthToken    string
		ServiceAuthToken string
		CollectionID     string
		ImageID          string
	}
	lockImageClientMockGetImage.RLock()
	calls = mock.calls.GetImage
	lockImageClientMockGetImage.RUnlock()
	return calls
}

var (
	lockRenderClientMockDo sync.RWMutex
)

// Ensure, that RenderClientMock does implement RenderClient.
// If this is not the case, regenerate this file with moq.
var _ RenderClient = &RenderClientMock{}

// RenderClientMock is a mock implementation of RenderClient.
//
//     func TestSomethingThatUsesRenderClient(t *testing.T) {
//
//         // make and configure a mocked RenderClient
//         mockedRenderClient := &RenderClientMock{
//             DoFunc: func(in1 string, in2 []byte) ([]byte, error) {
// 	               panic("mock out the Do method")
//             },
//         }
//
//         // use mockedRenderClient in code that requires RenderClient
//         // and then make assertions.
//
//     }
type RenderClientMock struct {
	// DoFunc mocks the Do method.
	DoFunc func(in1 string, in2 []byte) ([]byte, error)

	// calls tracks calls to the methods.
	calls struct {
		// Do holds details about calls to the Do method.
		Do []struct {
			// In1 is the in1 argument value.
			In1 string
			// In2 is the in2 argument value.
			In2 []byte
		}
	}
}

// Do calls DoFunc.
func (mock *RenderClientMock) Do(in1 string, in2 []byte) ([]byte, error) {
	if mock.DoFunc == nil {
		panic("RenderClientMock.DoFunc: method is nil but RenderClient.Do was just called")
	}
	callInfo := struct {
		In1 string
		In2 []byte
	}{
		In1: in1,
		In2: in2,
	}
	lockRenderClientMockDo.Lock()
	mock.calls.Do = append(mock.calls.Do, callInfo)
	lockRenderClientMockDo.Unlock()
	return mock.DoFunc(in1, in2)
}

// DoCalls gets all the calls that were made to Do.
// Check the length with:
//     len(mockedRenderClient.DoCalls())
func (mock *RenderClientMock) DoCalls() []struct {
	In1 string
	In2 []byte
} {
	var calls []struct {
		In1 string
		In2 []byte
	}
	lockRenderClientMockDo.RLock()
	calls = mock.calls.Do
	lockRenderClientMockDo.RUnlock()
	return calls
}
