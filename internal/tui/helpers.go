package tui

import "encoding/json"

func IndexOf[T comparable](arr []T, target T) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func MarshalJSONForPlaceholder(v any) string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return ""
	}
	if len(jsonData) == 0 || string(jsonData) == "null" {
		return ""
	}
	return string(jsonData)
}
