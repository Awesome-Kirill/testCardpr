package testCardpr

func CopyMap(original map[string]interface{}) map[string]interface{} {
	targetMap := make(map[string]interface{})

	// Copy from the original map to the target map
	for key, value := range original {
		targetMap[key] = value
	}

	return targetMap
}
