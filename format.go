package golog

type JSONformat struct {
	JSON map[Logger]string
}

func SetJSONformat(data *JSONformat) *JSONformat {
	return data
}
