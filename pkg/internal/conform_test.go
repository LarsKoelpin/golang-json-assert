package internal

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConform(t *testing.T) {
	Convey("Conform", t, func() {
		Convey("Shall conform", func() {
			actual := make(map[string]interface{}, 0)
			expected := make(map[string]interface{}, 0)
			actual["one"] = 12
			actual["two"] = 12
			expected["one"] = 12
			So(Conform(actual, expected), ShouldEqual, true)
		})

		Convey("Shall not conform when value has different type", func() {
			actual := make(map[string]interface{}, 0)
			expected := make(map[string]interface{}, 0)
			actual["one"] = 12
			actual["two"] = 12
			expected["one"] = "12"
			So(Conform(actual, expected), ShouldEqual, false)
		})

		Convey("Shall not conform when value is different", func() {
			actual := make(map[string]interface{}, 0)
			expected := make(map[string]interface{}, 0)
			actual["one"] = 12
			actual["two"] = 12
			expected["one"] = 15
			So(Conform(actual, expected), ShouldEqual, false) // 15 != 12
		})

		Convey("Shall not conform when actual has not every key of expectation", func() {
			actual := make(map[string]interface{}, 0)
			expected := make(map[string]interface{}, 0)
			actual["one"] = 12
			actual["two"] = 12
			expected["one"] = 12
			expected["uneedthisToo"] = 12
			So(Conform(actual, expected), ShouldEqual, false)
		})
	})
}
