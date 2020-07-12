package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	Convey("Unmarshal", t, func() {
		Convey("Unmarshal JSON shall unmarshal string to map", func() {
			aMap := UmmarshalToMap(`{"name": "lars", "age": 12}`)
			_, hasName := aMap["name"]
			_, hasAge := aMap["age"]
			So(hasName, ShouldEqual, true)
			So(hasAge, ShouldEqual, true)
		})

		Convey("Unmarshal JSON shall unmarshal list to map", func() {
			aMap := UmmarshalToMap(`{"name": "lars", "age": 12, "friends": []}`)
			_, hasName := aMap["name"]
			_, hasAge := aMap["age"]
			_, hasFriends := aMap["friends"]
			So(hasName, ShouldEqual, true)
			So(hasAge, ShouldEqual, true)
			So(hasFriends, ShouldEqual, true)
		})
	})

	Convey("Expect#ToEqual", t, func() {
		Convey("It shall be true, if all key and values are the same", func() {
			actual := `{"name": "lars", "age": 12}`
			expected := `{"name": "lars", "age": 12}`
			So(Expect(actual).ToEqual(expected), ShouldEqual, true)
		})

		Convey("It shall be false, if a key is missing", func() {
			actual := `{"name": "lars", "age": 12}`
			expected := `{"age": 12}`
			So(Expect(actual).ToEqual(expected), ShouldEqual, false)
		})

		Convey("It shall be false, if a array value is diffrent", func() {
			actual := `{"name": "lars", "age": 12, "friends": ["lars"]}`
			expected := `{"name": "lars", "age": 12, "friends": ["eric"]}`
			So(Expect(actual).ToEqual(expected), ShouldEqual, false)
		})
	})

	Convey("Expect#ToConform", t, func() {
		Convey("It shall be true if json conforms", func() {
			actual := `{"name": "lars", "age": 12}`
			expected := `{"age": 12}`
			So(Expect(actual).ToConform(expected), ShouldEqual, true)
		})

		Convey("It shall be false if json value does not conform", func() {
			actual := `{"name": "lars", "age": 12}`
			expected := `{"age": 13}`
			So(Expect(actual).ToConform(expected), ShouldEqual, false)
		})

		Convey("It shall be false if json keys do not conform", func() {
			actual := `{"name": "lars", "age": 12}`
			expected := `{"street": 13}`
			So(Expect(actual).ToConform(expected), ShouldEqual, false)
		})
	})
}
