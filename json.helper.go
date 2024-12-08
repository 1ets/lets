package lets

import (
	"encoding/json"
)

// Encode json from object to JSON and beautify the output.
func ToJson(data interface{}) string {
	jsonResult, _ := json.Marshal(data)

	return string(jsonResult)
}

// Encode json from object to JSON and beautify the output.
func ToJsonIndent(data interface{}) string {
	jsonResult, _ := json.MarshalIndent(data, "", " ")

	return string(jsonResult)
}
