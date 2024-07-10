package handlers

import "reflect"

type defaultResponseTemplate[T any] struct {
	Success bool   `json:"success"` // true or false
	Error   string `json:"error"`   // success or failure message
	Data    []T    `json:"data"`    // list of data based on type
}

// Create Factory
// NewErrorResponse creates a new error response with an empty data slice.
func NewErrorResponse[T any](errMsg string) defaultResponseTemplate[T] {
	return defaultResponseTemplate[T]{
		Success: false,
		Error:   errMsg,
		Data:    []T{},
	}
}

// NewSuccessResponse creates a new success response with data.
func NewSuccessResponse[T any](data interface{}) defaultResponseTemplate[T] {
	var responseData []T

	// Use reflection to check the type of data
	val := reflect.ValueOf(data)
	switch val.Kind() {
	case reflect.Slice:
		// If the data is already a slice, convert it directly
		responseData = data.([]T)
	default:
		// Otherwise, wrap it into a slice
		responseData = []T{data.(T)}
	}

	return defaultResponseTemplate[T]{
		Success: true,
		Error:   "",
		Data:    responseData,
	}
}
