package slices

func Contains(sliceSearch []interface{}, valueSearch interface{}) bool {

	for _, value := range sliceSearch {
		if value == valueSearch {
			return true
		}
	}
	return false

}
