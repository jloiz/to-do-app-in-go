package helpers

func FindInArray(chkVal interface{}, chkArr []interface{}) (int) {
	// Checks if a value is in an array and returns the position if present
	// Returns -1 if the value is not found in the array
	for i, v := range chkArr {
		if v == chkVal {
			return i
		}
	}
	return -1
}
