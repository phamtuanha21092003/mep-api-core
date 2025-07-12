package base

import "encoding/json"

func convertToMap(input any) (map[string]interface{}, bool) {
	jsonData, err := json.Marshal(input)
	if err != nil {
		return nil, false
	}

	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, false
	}
	return result, true
}
