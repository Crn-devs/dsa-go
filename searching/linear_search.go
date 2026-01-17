package searching

func LinearSearch(arr []int, target int) int {

	if len(arr) <= 0 {
		return -1
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

// indexof returns the index of the val if the val is not present it returns -1
func IndexOfString(arr []string, val string) int {

	if len(arr) <= 0 {
		return -1
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return i
		}
	}
	return -1
}

// contains checks if val is in the array if val is found true is returned if not false is returned
func Contains(arr []string, val string) bool {
	if len(arr) <= 0 {
		return false
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return true
		}
	}
	return false
}
