// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package homepage

import (
	"context"
	"github.com/ONSdigital/dp-frontend-homepage-controller/model"
	"github.com/ONSdigital/dp-topic-api/models"
	"sync"
	"time"
)

var (
	lockClienterMockAddNavigationCache    sync.RWMutex
	lockClienterMockClose                 sync.RWMutex
	lockClienterMockGetHomePage           sync.RWMutex
	lockClienterMockGetNavigationData     sync.RWMutex
	lockClienterMockStartBackgroundUpdate sync.RWMutex
)

// Ensure, that ClienterMock does implement Clienter.
// If this is not the case, regenerate this file with moq.
var _ Clienter = &ClienterMock{}

// ClienterMock is a mock implementation of Clienter.
//
//     func TestSomethingThatUsesClienter(t *testing.T) {
//
//         // make and configure a mocked Clienter
//         mockedClienter := &ClienterMock{
//             AddNavigationCacheFunc: func(ctx context.Context, updateInterval time.Duration) error {
// 	               panic("mock out the AddNavigationCache method")
//             },
//             CloseFunc: func()  {
// 	               panic("mock out the Close method")
//             },
//             GetHomePageFunc: func(ctx context.Context, userAccessToken string, collectionID string, lang string) (*model.HomepageData, error) {
// 	               panic("mock out the GetHomePage method")
//             },
//             GetNavigationDataFunc: func(ctx context.Context, lang string) (*models.Navigation, error) {
// 	               panic("mock out the GetNavigationData method")
//             },
//             StartBackgroundUpdateFunc: func(ctx context.Context, errorChannel chan error)  {
// 	               panic("mock out the StartBackgroundUpdate method")
//             },
//         }
//
//         // use mockedClienter in code that requires Clienter
//         // and then make assertions.
//
//     }
type ClienterMock struct {
	// AddNavigationCacheFunc mocks the AddNavigationCache method.
	AddNavigationCacheFunc func(ctx context.Context, updateInterval time.Duration) error

	// CloseFunc mocks the Close method.
	CloseFunc func()

	// GetHomePageFunc mocks the GetHomePage method.
	GetHomePageFunc func(ctx context.Context, userAccessToken string, collectionID string, lang string) (*model.HomepageData, error)

	// GetNavigationDataFunc mocks the GetNavigationData method.
	GetNavigationDataFunc func(ctx context.Context, lang string) (*models.Navigation, error)

	// StartBackgroundUpdateFunc mocks the StartBackgroundUpdate method.
	StartBackgroundUpdateFunc func(ctx context.Context, errorChannel chan error)

	// calls tracks calls to the methods.
	calls struct {
		// AddNavigationCache holds details about calls to the AddNavigationCache method.
		AddNavigationCache []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UpdateInterval is the updateInterval argument value.
			UpdateInterval time.Duration
		}
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// GetHomePage holds details about calls to the GetHomePage method.
		GetHomePage []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserAccessToken is the userAccessToken argument value.
			UserAccessToken string
			// CollectionID is the collectionID argument value.
			CollectionID string
			// Lang is the lang argument value.
			Lang string
		}
		// GetNavigationData holds details about calls to the GetNavigationData method.
		GetNavigationData []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Lang is the lang argument value.
			Lang string
		}
		// StartBackgroundUpdate holds details about calls to the StartBackgroundUpdate method.
		StartBackgroundUpdate []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ErrorChannel is the errorChannel argument value.
			ErrorChannel chan error
		}
	}
}

// AddNavigationCache calls AddNavigationCacheFunc.
func (mock *ClienterMock) AddNavigationCache(ctx context.Context, updateInterval time.Duration) error {
	if mock.AddNavigationCacheFunc == nil {
		panic("ClienterMock.AddNavigationCacheFunc: method is nil but Clienter.AddNavigationCache was just called")
	}
	callInfo := struct {
		Ctx            context.Context
		UpdateInterval time.Duration
	}{
		Ctx:            ctx,
		UpdateInterval: updateInterval,
	}
	lockClienterMockAddNavigationCache.Lock()
	mock.calls.AddNavigationCache = append(mock.calls.AddNavigationCache, callInfo)
	lockClienterMockAddNavigationCache.Unlock()
	return mock.AddNavigationCacheFunc(ctx, updateInterval)
}

// AddNavigationCacheCalls gets all the calls that were made to AddNavigationCache.
// Check the length with:
//     len(mockedClienter.AddNavigationCacheCalls())
func (mock *ClienterMock) AddNavigationCacheCalls() []struct {
	Ctx            context.Context
	UpdateInterval time.Duration
} {
	var calls []struct {
		Ctx            context.Context
		UpdateInterval time.Duration
	}
	lockClienterMockAddNavigationCache.RLock()
	calls = mock.calls.AddNavigationCache
	lockClienterMockAddNavigationCache.RUnlock()
	return calls
}

// Close calls CloseFunc.
func (mock *ClienterMock) Close() {
	if mock.CloseFunc == nil {
		panic("ClienterMock.CloseFunc: method is nil but Clienter.Close was just called")
	}
	callInfo := struct {
	}{}
	lockClienterMockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	lockClienterMockClose.Unlock()
	mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedClienter.CloseCalls())
func (mock *ClienterMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	lockClienterMockClose.RLock()
	calls = mock.calls.Close
	lockClienterMockClose.RUnlock()
	return calls
}

// GetHomePage calls GetHomePageFunc.
func (mock *ClienterMock) GetHomePage(ctx context.Context, userAccessToken string, collectionID string, lang string) (*model.HomepageData, error) {
	if mock.GetHomePageFunc == nil {
		panic("ClienterMock.GetHomePageFunc: method is nil but Clienter.GetHomePage was just called")
	}
	callInfo := struct {
		Ctx             context.Context
		UserAccessToken string
		CollectionID    string
		Lang            string
	}{
		Ctx:             ctx,
		UserAccessToken: userAccessToken,
		CollectionID:    collectionID,
		Lang:            lang,
	}
	lockClienterMockGetHomePage.Lock()
	mock.calls.GetHomePage = append(mock.calls.GetHomePage, callInfo)
	lockClienterMockGetHomePage.Unlock()
	return mock.GetHomePageFunc(ctx, userAccessToken, collectionID, lang)
}

// GetHomePageCalls gets all the calls that were made to GetHomePage.
// Check the length with:
//     len(mockedClienter.GetHomePageCalls())
func (mock *ClienterMock) GetHomePageCalls() []struct {
	Ctx             context.Context
	UserAccessToken string
	CollectionID    string
	Lang            string
} {
	var calls []struct {
		Ctx             context.Context
		UserAccessToken string
		CollectionID    string
		Lang            string
	}
	lockClienterMockGetHomePage.RLock()
	calls = mock.calls.GetHomePage
	lockClienterMockGetHomePage.RUnlock()
	return calls
}

// GetNavigationData calls GetNavigationDataFunc.
func (mock *ClienterMock) GetNavigationData(ctx context.Context, lang string) (*models.Navigation, error) {
	if mock.GetNavigationDataFunc == nil {
		panic("ClienterMock.GetNavigationDataFunc: method is nil but Clienter.GetNavigationData was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Lang string
	}{
		Ctx:  ctx,
		Lang: lang,
	}
	lockClienterMockGetNavigationData.Lock()
	mock.calls.GetNavigationData = append(mock.calls.GetNavigationData, callInfo)
	lockClienterMockGetNavigationData.Unlock()
	return mock.GetNavigationDataFunc(ctx, lang)
}

// GetNavigationDataCalls gets all the calls that were made to GetNavigationData.
// Check the length with:
//     len(mockedClienter.GetNavigationDataCalls())
func (mock *ClienterMock) GetNavigationDataCalls() []struct {
	Ctx  context.Context
	Lang string
} {
	var calls []struct {
		Ctx  context.Context
		Lang string
	}
	lockClienterMockGetNavigationData.RLock()
	calls = mock.calls.GetNavigationData
	lockClienterMockGetNavigationData.RUnlock()
	return calls
}

// StartBackgroundUpdate calls StartBackgroundUpdateFunc.
func (mock *ClienterMock) StartBackgroundUpdate(ctx context.Context, errorChannel chan error) {
	if mock.StartBackgroundUpdateFunc == nil {
		panic("ClienterMock.StartBackgroundUpdateFunc: method is nil but Clienter.StartBackgroundUpdate was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		ErrorChannel chan error
	}{
		Ctx:          ctx,
		ErrorChannel: errorChannel,
	}
	lockClienterMockStartBackgroundUpdate.Lock()
	mock.calls.StartBackgroundUpdate = append(mock.calls.StartBackgroundUpdate, callInfo)
	lockClienterMockStartBackgroundUpdate.Unlock()
	mock.StartBackgroundUpdateFunc(ctx, errorChannel)
}

// StartBackgroundUpdateCalls gets all the calls that were made to StartBackgroundUpdate.
// Check the length with:
//     len(mockedClienter.StartBackgroundUpdateCalls())
func (mock *ClienterMock) StartBackgroundUpdateCalls() []struct {
	Ctx          context.Context
	ErrorChannel chan error
} {
	var calls []struct {
		Ctx          context.Context
		ErrorChannel chan error
	}
	lockClienterMockStartBackgroundUpdate.RLock()
	calls = mock.calls.StartBackgroundUpdate
	lockClienterMockStartBackgroundUpdate.RUnlock()
	return calls
}