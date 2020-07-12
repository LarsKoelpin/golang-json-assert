package internal

import (
	"log"
	"reflect"
)

func Conform(actual interface{}, expectation interface{}) bool {
	if IsArray(actual) && IsArray(expectation) {
		actualArr, _ := actual.([]interface{})
		expectationArr, _ := expectation.([]interface{})
		return ConformArray(actualArr, expectationArr)
	}

	if IsObject(actual) && IsObject(expectation) {
		actualObj, _ := actual.(map[string]interface{})
		expectedObj, _ := expectation.(map[string]interface{})
		return ConformObject(actualObj, expectedObj)
	}

	panic("UNKNOWN TYPE")
}

func ConformArray(actual []interface{}, expected []interface{}) bool {
	actualLen := len(actual)

	for i := 0; i < actualLen; i++ {
		aValue := actual[i]
		eValue := expected[i]

		if !HaveSameType(aValue, eValue) {
			return false
		}

		if IsObject(aValue) && IsObject(eValue) {
			aValueObject, _ := aValue.(map[string]interface{})
			eValueObject, _ := eValue.(map[string]interface{})
			if !ConformObject(aValueObject, eValueObject) {
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
			if !ConformArray(aValueArray, eValueArray) {
				return false
			}
		}
	}

	return true
}

// {name: "aa"} == {name: "bb"}
func ConformObject(actual map[string]interface{}, expected map[string]interface{}) bool {
	actualType := reflect.TypeOf(actual)
	typesEqual := actualType == reflect.TypeOf(expected)

	if !typesEqual {
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
			if !ConformObject(actualChildObject, expectedChildObject) {
				return false
			}
		}

		if IsArray(actualValue) && IsArray(expectedValue) {
			actualChildObject, _ := actualValue.([]interface{})
			expectedChildObject, _ := expectedValue.([]interface{})
			if !ConformArray(actualChildObject, expectedChildObject) {
				return false
			}
		}
	}
	return true
}
