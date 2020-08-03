package goconvey

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	convey.Convey("两数相加", t, func() {
		convey.So(Add(1, 2), convey.ShouldEqual, 3)
		convey.So(Add(2, 3), convey.ShouldEqual, 5)
		convey.So(Add(-1, 2), convey.ShouldEqual, 1)

	})
}
