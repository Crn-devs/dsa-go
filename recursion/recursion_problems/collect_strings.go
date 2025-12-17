package recursionproblems

/**
collectStrings

Write a function called collectStrings which accepts an object
and returns an array of all the values in the object that have a typeof string

**/

func CollectStrings(obj map[string]any) []string {

	var strings []string
	// if len(obj) == 0 {
	// 	return strings
	// }

	for _, val := range obj {
		switch v := val.(type) {
		case string:
			strings = append(strings, v)
		case map[string]any:
			strings = append(strings, CollectStrings(v)...)
		default:
			continue
		}
	}
	return strings
}
