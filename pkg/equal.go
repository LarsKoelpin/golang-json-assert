package main

import "reflect"

func StrictEqualArray(actual []interface{}, expected []interface{}) bool {

	

	return false
}

func GetTypeArray(arr interface{}) reflect.Type {
	return reflect.TypeOf(arr).Elem()
}

func StrictEqual(actual map[string]interface{}, expectation map[string]interface{}) bool {

	actualLen := len(actual)
	expectedLen := len(expectation)

	if actualLen != expectedLen {
		return false
	}

	for key, _ := range expectation {
		actualValue, hasActualKey := actual[key]
		expectedValue, hasExpectedKey := expectation[key]

		typesEqual := reflect.TypeOf(actualValue) == reflect.TypeOf(expectedValue)

		if !typesEqual {
			return false
		}

		if !hasActualKey && !hasExpectedKey {
			return false
		}

		if actualValue != expectedValue {
			return false
		}
	}
	return true
}

func KeyEquality(actual map[string]interface{}, expectation map[string]interface{}) bool {
	for key, _ := range expectation {
		_, hasActualKey := actual[key]
		if !hasActualKey {
			return false
		}
	}
	return true
}

func Conform(actual map[string]interface{}, expectation map[string]interface{}) bool {
	for key, value := range expectation {
		actualValue, hasCorrespondingActual := actual[key]

		if !hasCorrespondingActual {
			return false
		}

		if hasCorrespondingActual {
			if actualValue != value {
				return false
			}
		}
	}
	return true
}

func ConformKeys(actual map[string]interface{}, expectation map[string]interface{}) bool {
	for key, _ := range expectation {
		_, hasCorrespondingActual := actual[key]

		if !hasCorrespondingActual {
			return false
		}
	}
	return true
}

// Shallow Equal
