package internal

import "reflect"

func HaveSameType(a interface{}, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

func IsObject(a interface{}) bool {
	actualValueType := reflect.TypeOf(a)
	return actualValueType.Kind() == reflect.Map
}

func IsPrimitive(a interface{}) bool {
	actualValueType := reflect.TypeOf(a)
	return actualValueType.Kind() == reflect.Int || actualValueType.Kind() == reflect.String || actualValueType.Kind() == reflect.Float32 || actualValueType.Kind() == reflect.Float64
}

func IsArray(a interface{}) bool {
	actualValueType := reflect.TypeOf(a)
	return actualValueType.Kind() == reflect.Slice || actualValueType.Kind() == reflect.Array
}
