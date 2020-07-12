package main

import (
	"encoding/json"
  "log"
  "strings"
)

type ActualJson struct {
	jsonMap map[string]interface{}
	jsonArray []interface{}
}

func Expect(str string) ActualJson {
  if strings.HasPrefix(str, "[") {
    arr := UnmarshalToArray(str);
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
      return StrictEqualArray(a.jsonArray, expect);
    }
    log.Print("Cannot compare array with object")
    return false
  }
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

func UnmarshalToArray(src string) []interface{} {
  var m []interface{}
  err := json.Unmarshal([]byte(src), &m)
  if err != nil {
    panic(err)
  }
  return m
}

