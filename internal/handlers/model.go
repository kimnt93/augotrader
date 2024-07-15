package handlers

import (
	"reflect"
)

type defaultResponseTemplate[T any] struct {
	Success bool   `json:"success"` // true or false
	Error   string `json:"error"`   // success or failure message
	Data    []T    `json:"data"`    // list of data based on type, if null, return empty list
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
	// Check if data is nil, return an empty slice if so
	if data == nil {
		return defaultResponseTemplate[T]{
			Success: true,
			Error:   "",
			Data:    []T{},
		}
	}

	// Use reflection to check the type of data
	val := reflect.ValueOf(data)
	switch val.Kind() {
	case reflect.Slice:
		if val.Len() == 0 {
			// If the slice is empty, return an empty response
			return defaultResponseTemplate[T]{
				Success: true,
				Error:   "",
				Data:    []T{},
			}
		}

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
