package mapper

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitMapper(t *testing.T) {
	ctx := context.Background()

	Convey("test homepage mapping works", t, func() {
		page := Homepage(ctx)

		So(page.Type, ShouldEqual, "homepage")
	})
}
