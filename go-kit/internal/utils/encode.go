package utils

import (
	"context"
	"encoding/json"
	"net/http"
)

func EncodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func EncodeHTTPResponseWithData(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	data := map[string]interface{}{
		"data": response,
	}
	return json.NewEncoder(w).Encode(data)
}

// EncodeLegacyError is used to mirror MPG1's error response format
func EncodeLegacyError(_ context.Context, err error, w http.ResponseWriter) {
	errResponse := new(Error)
	errResponse, ok := err.(*Error)
	if !ok {
		errResponse = ErrBadRequest
	}

	code := errResponse.Code / 100

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errResponse)
}
