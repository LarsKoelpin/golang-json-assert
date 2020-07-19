package internal

import (
	"log"
	"reflect"
)

func ConformWithWarning(actual interface{}, expectation interface{}, shallwarn bool) bool {
	if IsArray(actual) && IsArray(expectation) {
		actualArr, _ := actual.([]interface{})
		expectationArr, _ := expectation.([]interface{})
		return ConformArray(actualArr, expectationArr, shallwarn)
	}

	if IsObject(actual) && IsObject(expectation) {
		actualObj, _ := actual.(map[string]interface{})
		expectedObj, _ := expectation.(map[string]interface{})
		return ConformObject(actualObj, expectedObj, shallwarn)
	}

	panic("UNKNOWN TYPE")
}

func Conform(actual interface{}, expectation interface{}) bool {
	if IsArray(actual) && IsArray(expectation) {
		actualArr, _ := actual.([]interface{})
		expectationArr, _ := expectation.([]interface{})
		return ConformArray(actualArr, expectationArr, true)
	}

	if IsObject(actual) && IsObject(expectation) {
		actualObj, _ := actual.(map[string]interface{})
		expectedObj, _ := expectation.(map[string]interface{})
		return ConformObject(actualObj, expectedObj, true)
	}

	panic("UNKNOWN TYPE")
}

func ConformArray(actual []interface{}, expected []interface{}, shallWarn bool) bool {
	expectedLen := len(expected)
	seen := make(map[int]bool)
	for i := 0; i < expectedLen; i++ {
		eValue := expected[i]
		if !findAnyConformObject(eValue, actual, seen) && seen[i] == false {
			seen[i] = true
			return false
		}
	}
	return true
}

// {name: "aa"} == {name: "bb"}
func ConformObject(actual map[string]interface{}, expected map[string]interface{}, warn bool) bool {
	actualType := reflect.TypeOf(actual)
	typesEqual := actualType == reflect.TypeOf(expected)

	if !typesEqual {
		return false
	}

	for key := range expected {
		actualValue, hasActualKey := actual[key]
		expectedValue, hasExpectedKey := expected[key]

		if hasExpectedKey && !hasActualKey {
			if warn {
				log.Print("Key difference at key", key)
			}
			return false
		}

		if !HaveSameType(actualValue, expectedValue) {
			if warn {
				log.Print("Type difference at key ", key)
			}
			return false
		}

		if IsPrimitive(actualValue) && IsPrimitive(expectedValue) {
			if !EqualPrimitive(actualValue, expectedValue) {
				if warn {
					log.Print("Value difference at key: ", key)
				}
				return false
			}
		}

		if IsObject(actualValue) && IsObject(expectedValue) {
			actualChildObject, _ := actualValue.(map[string]interface{})
			expectedChildObject, _ := expectedValue.(map[string]interface{})
			if !ConformObject(actualChildObject, expectedChildObject, warn) {
				if warn {
					log.Print("Object difference at key ", key)
				}
				return false
			}
		}

		if IsArray(actualValue) && IsArray(expectedValue) {
			actualChildObject, _ := actualValue.([]interface{})
			expectedChildObject, _ := expectedValue.([]interface{})
			if !ConformArray(actualChildObject, expectedChildObject, warn) {
				if warn {
					log.Print("Array difference at key ", key)
				}
				return false
			}
		}
	}
	return true
}

func findAnyConformObject(wanted interface{}, actual []interface{}, seen map[int]bool) bool {
	for i := 0; i < len(actual); i++ {
		if ConformWithWarning(actual[i], wanted, false) && !seen[i] {
			seen[i] = true
			return true
		}
	}
	log.Print("No conformed object found")
	return false
}
