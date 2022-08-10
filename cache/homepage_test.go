package cache

import (
	"context"
	"testing"
	"time"

	"github.com/ONSdigital/dp-frontend-homepage-controller/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewHomepageCache(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	Convey("Given a valid cache update interval which is greater than 0", t, func() {
		updateCacheInterval := 1 * time.Millisecond

		Convey("When NewHomepageCache is called", func() {
			testCache, err := NewHomepageCache(ctx, &updateCacheInterval)

			Convey("Then a homepage cache object should be successfully returned", func() {
				So(testCache, ShouldNotBeEmpty)

				Convey("And no error should be returned", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})

	Convey("Given no cache update interval (nil)", t, func() {
		Convey("When NewHomepageCache is called", func() {
			testCache, err := NewHomepageCache(ctx, nil)

			Convey("Then a cache object should be successfully returned", func() {
				So(testCache, ShouldNotBeEmpty)

				Convey("And no error should be returned", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})

	Convey("Given an invalid cache update interval which is less than or equal to 0", t, func() {
		updateCacheInterval := 0 * time.Second

		Convey("When NewHomepageCache is called", func() {
			testCache, err := NewHomepageCache(ctx, &updateCacheInterval)

			Convey("Then an error should be returned", func() {
				So(err, ShouldNotBeNil)

				Convey("And a nil cache object should be returned", func() {
					So(testCache, ShouldBeNil)
				})
			})
		})
	})
}

func TestHomepageAddUpdateFunc(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	Convey("Given an update function to update a topic", t, func() {
		mockHomepageCache, err := NewHomepageCache(ctx, nil)
		So(err, ShouldBeNil)

		homepageUpdateFunc := func() (*model.HomepageData, error) {
			return &model.HomepageData{
				ServiceMessage: "Test",
			}, nil
		}

		Convey("When AddUpdateFunc is called", func() {
			mockHomepageCache.AddUpdateFunc("test", homepageUpdateFunc)

			Convey("Then the update function is added to the cache", func() {
				So(mockHomepageCache.UpdateFuncs["test"], ShouldNotBeEmpty)
			})
		})
	})
}
