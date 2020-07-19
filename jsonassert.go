package golang_json_assert

import (
	"encoding/json"
	"log"
	"strings"

	. "github.com/larskoelpin/golang-json-assert/internal"
)

type ActualJson struct {
	jsonMap   map[string]interface{}
	jsonArray []interface{}
}

func Expect(str string) ActualJson {
	if strings.HasPrefix(str, "[") {
		arr := UnmarshalToArray(str)
		return ActualJson{
			jsonArray: arr,
		}
	}
	theJsonMap := UmmarshalToMap(str)
	return ActualJson{jsonMap: theJsonMap}
}

func (a ActualJson) ToEqual(expectation string) bool {

	if a.jsonArray != nil {
		if strings.HasPrefix(expectation, "[") {
			expect := UnmarshalToArray(expectation)
			return EqualArray(a.jsonArray, expect)
		}
		log.Print("Cannot compare array with object")
		return false
	}
	expectationJsonMap := UmmarshalToMap(expectation)
	return StrictEqual(a.jsonMap, expectationJsonMap)
}

func (a ActualJson) ToConform(expectation string) bool {
	if a.jsonArray != nil {
		if strings.HasPrefix(expectation, "[") {
			expect := UnmarshalToArray(expectation)
			return ConformArray(a.jsonArray, expect, true)
		}
		log.Print("Cannot compare array with object")
		return false
	}
	expectationJsonMap := UmmarshalToMap(expectation)
	return Conform(a.jsonMap, expectationJsonMap)
}

func UmmarshalToMap(src string) map[string]interface{} {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(src), &m)
	if err != nil {
		panic(err)
	}
	return m
}

func UnmarshalToArray(src string) []interface{} {
	var m []interface{}
	err := json.Unmarshal([]byte(src), &m)
	if err != nil {
		panic(err)
	}
	return m
}
