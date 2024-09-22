package golog

type JSONformat struct {
	JSON map[string][]any `json:"data"`
}

func SetJSONformat(data *JSONformat) *JSONformat {
	return data
}
