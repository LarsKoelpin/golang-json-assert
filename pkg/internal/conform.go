package internal

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
