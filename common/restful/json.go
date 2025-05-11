package restful

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	contentTypeName = "Content-Type"
	contentTypeJSON = "application/json"
)

func WriteJSONResponse[T any](w http.ResponseWriter, payload T) {
	data, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(contentTypeName, contentTypeJSON)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func ReadRequest[T any](r *http.Request) (*T, error) {
	// read content-type header and decide which deserializer use
	contentType := r.Header.Get(contentTypeName)
	switch contentType {
	case contentTypeJSON:
		return readJSONRequest[T](r)
	default:
		return nil, fmt.Errorf("invalid Content-Type: %s", contentType)
	}
}

func readJSONRequest[T any](r *http.Request) (*T, error) {
	var payload T
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}
