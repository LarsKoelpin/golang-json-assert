package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEquality(t *testing.T) {
	Convey("StrictEqual", t, func() {

		Convey("Shall succeed simple StrictEquality", func() {
			actual := make(map[string]interface{}, 0)
			expected := make(map[string]interface{}, 0)
			actual["hello"] = 12
			expected["hello"] = 12
			So(StrictEqual(actual, expected), ShouldEqual, true)
		})

		Convey("Shall fail if types do not match", func() {
			actual := make(map[string]interface{}, 0)
			expected := make(map[string]interface{}, 0)
			actual["hello"] = 12
			expected["hello"] = "a"
			So(StrictEqual(actual, expected), ShouldEqual, false)
		})

		Convey("Shall fail simple StrictEquality for diffrent keys", func() {
			actual := make(map[string]interface{}, 0)
			expected := make(map[string]interface{}, 0)
			actual["hello"] = 12
			expected["nothello"] = 12
			So(StrictEqual(actual, expected), ShouldEqual, false)
		})

		Convey("Shall fail simple StrictEquality for different values", func() {
			actual := make(map[string]interface{}, 0)
			expected := make(map[string]interface{}, 0)
			actual["hello"] = 12
			expected["hello"] = 13
			So(StrictEqual(actual, expected), ShouldEqual, false)
		})

	})

	Convey("KeyEqual", t, func() {
		Convey("Shall succeed simple KeyEquality", func() {
			actual := make(map[string]interface{}, 0)
			expected := make(map[string]interface{}, 0)
			actual["hello"] = 12
			expected["hello"] = 15
			So(KeyEquality(actual, expected), ShouldEqual, true)
		})

		Convey("Shall fail simple KeyEquality", func() {
			actual := make(map[string]interface{}, 0)
			expected := make(map[string]interface{}, 0)
			actual["hello"] = 12
			expected["xd"] = 15
			So(KeyEquality(actual, expected), ShouldEqual, false)
		})
	})

	Convey("Conform", t, func() {
		Convey("Shall conform", func() {
			actual := make(map[string]interface{}, 0)
			expected := make(map[string]interface{}, 0)
			actual["one"] = 12
			actual["two"] = 12
			expected["one"] = 12
			So(Conform(actual, expected), ShouldEqual, true) // Onforms, as Expcted only contians key "one"
		})

		Convey("Shall not conform when value is diffrent types", func() {
			actual := make(map[string]interface{}, 0)
			expected := make(map[string]interface{}, 0)
			actual["one"] = 12
			actual["two"] = 12
			expected["one"] = "12"
			So(Conform(actual, expected), ShouldEqual, false) // Onforms, as Expcted only contians key "one"
		})

		Convey("Shall not conform when value is diffrent", func() {
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

	Convey("ConformKeys", t, func() {
		Convey("Shall conform with same value", func() {
			actual := make(map[string]interface{}, 0)
			expected := make(map[string]interface{}, 0)
			actual["one"] = 12
			actual["two"] = 12
			expected["one"] = 12
			So(ConformKeys(actual, expected), ShouldEqual, true) // Onforms, as Expcted only contians key "one"
		})

		Convey("Shall conform Keys with diffret value", func() {
			actual := make(map[string]interface{}, 0)
			expected := make(map[string]interface{}, 0)
			actual["one"] = 12
			actual["two"] = 12
			expected["one"] = 1255
			So(ConformKeys(actual, expected), ShouldEqual, true)
		})
	})
}
