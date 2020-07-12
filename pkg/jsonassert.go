package main

import (
	"encoding/json"
)

type ActualJson struct {
	jsonMap map[string]interface{}
}

func Expect(str string) ActualJson {
	theJsonMap := UmmarshalToMap(str)
	return ActualJson{jsonMap: theJsonMap}
}

func (a ActualJson) ToEqual(expectation string) bool {
	expectationJsonMap := UmmarshalToMap(expectation)
	return StrictEqual(a.jsonMap, expectationJsonMap)
}

func (a ActualJson) ToConform(expectation string) bool {
	expectationJsonMap := UmmarshalToMap(expectation)
	return Conform(a.jsonMap, expectationJsonMap)
}

type Jsontype struct {
	jsonType   string // Object, Primitive, Array
	objValue   map[string]interface{}
	atomValue  interface{}
	arrayValue []interface{}
}

func UmmarshalToMap(src string) map[string]interface{} {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(src), &m)
	if err != nil {
		panic(err)
	}
	return m
}
