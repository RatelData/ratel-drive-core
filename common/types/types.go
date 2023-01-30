package types

// HTTP json result wrapper
type JSONResult struct {
	Data string `json:"data"`
}

// HTTP error result wrapper
type ErrorResult struct {
	Error string `json:"error"`
}
