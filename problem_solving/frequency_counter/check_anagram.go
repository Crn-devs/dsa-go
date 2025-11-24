package main

func CheckAnagram(str1, str2 string) bool {

	// if lengths arent the same then automatically false
	if len(str1) != len(str2) {
		return false
	}

	// create maps to hold string values and counts
	frequency1 := make(map[string]int, len(str1))
	frequency2 := make(map[string]int, len(str2))

	// write the values to the map and increment the count for each letter
	for _, r := range str1 {
		frequency1[string(r)] = frequency1[string(r)] + 1
	}

	// same again for map 2
	for _, r := range str2 {
		frequency2[string(r)] = frequency2[string(r)] + 1
	}

	// loop over one of the frequency counters and check to make sure that the counts match
	// if the counts of each letter in the map do not match then return false
	for key, count := range frequency1 {
		if frequency2[key] != count {
			return false
		}
	}
	/**
	return true - all letters and counts match so return true
		as the strings contain the same letters and the same amount of letters
	**/

	return true

}
