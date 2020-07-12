package main

import (
	"log"
	"reflect"
)

// 12 == 12
func EqualPrimitive(actual interface{}, expected interface{}) bool {
	return actual == expected
}

func EqualArray(actual []interface{}, expected []interface{}) bool {
	actualLen := len(actual)
	expectedLen := len(expected)
	if actualLen != expectedLen {
		return false
	}

	for i := 0; i < actualLen; i++ {
		aValue := actual[i]
		eValue := expected[i]

		if !HaveSameType(aValue, eValue) {
			return false
		}

		if IsObject(aValue) && IsObject(eValue) {
			aValueObject, _ := aValue.(map[string]interface{})
			eValueObject, _ := eValue.(map[string]interface{})
			if !EqualObject(aValueObject, eValueObject) {
				return false
			}
		}

		if IsPrimitive(aValue) && IsPrimitive(eValue) {
			if !EqualPrimitive(aValue, eValue) {
				return false
			}
		}

		if IsArray(aValue) && IsArray(eValue) {
			aValueArray := aValue.([]interface{})
			eValueArray := aValue.([]interface{})
			if !EqualArray(aValueArray, eValueArray) {
				return false
			}
		}
	}

	return true
}

// {name: "aa"} == {name: "bb"}
func EqualObject(actual map[string]interface{}, expected map[string]interface{}) bool {
	actualType := reflect.TypeOf(actual)
	typesEqual := actualType == reflect.TypeOf(expected)

	if !typesEqual {
		return false
	}

	if len(expected) != len(actual) {
		return false
	}

	for key := range expected {
		actualValue, hasActualKey := actual[key]
		expectedValue, hasExpectedKey := expected[key]

		if hasExpectedKey && !hasActualKey {
			log.Print("Key difference")
			return false
		}

		if !HaveSameType(actualValue, expectedValue) {
			log.Print("Type difference")
			return false
		}

		if IsPrimitive(actualValue) && IsPrimitive(expectedValue) {
			if !EqualPrimitive(actualValue, expectedValue) {
				return false
			}
		}

		if IsObject(actualValue) && IsObject(expectedValue) {
			actualChildObject, _ := actualValue.(map[string]interface{})
			expectedChildObject, _ := expectedValue.(map[string]interface{})
			if !EqualObject(actualChildObject, expectedChildObject) {
				return false
			}
		}

		if IsArray(actualValue) && IsArray(expectedValue) {
			actualChildObject, _ := actualValue.([]interface{})
			expectedChildObject, _ := expectedValue.([]interface{})
			if !EqualArray(actualChildObject, expectedChildObject) {
				return false
			}
		}
	}
	return true
}

func StrictEqual(actual interface{}, expectation interface{}) bool {
	if IsArray(actual) && IsArray(expectation) {
		actualArr, _ := actual.([]interface{})
		expectationArr, _ := expectation.([]interface{})
		return EqualArray(actualArr, expectationArr)
	}

	if IsObject(actual) && IsObject(expectation) {
		actualObj, _ := actual.(map[string]interface{})
		expectedObj, _ := expectation.(map[string]interface{})
		return EqualObject(actualObj, expectedObj)
	}

	panic("UNKNOWN TYPE")
}
