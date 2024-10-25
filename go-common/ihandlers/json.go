package ihandlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/what-da-flac/wtf/go-common/exceptions"
)

// ReadJSON parses a reader into a target struct.
func ReadJSON(r io.Reader, v interface{}) error {
	if r == nil {
		return fmt.Errorf("empty reader")
	}
	return json.NewDecoder(r).Decode(v)
}

// writeJSON writes to a writer the target struct.
func writeJSON(w io.Writer, v interface{}) error {
	if w == nil {
		return fmt.Errorf("empty writer")
	}
	return json.NewEncoder(w).Encode(v)
}

func WriteResponse(w http.ResponseWriter, statusCode int, res interface{}, e error) {
	const (
		contentTypeKey = "Content-Type"
		jsonMime       = "application/json"
	)
	var errMsg string
	if e != nil {
		// check if we can infer the status code automatically, based on the error
		httpErr := exceptions.NewHTTPError(e)
		if val := httpErr.StatusCode(); val != 0 {
			statusCode = val
		}
		errMsg = httpErr.Error()
	}
	if res == nil && e != nil {
		res = struct {
			Error string `json:"error"`
		}{
			Error: errMsg,
		}
	}
	w.WriteHeader(statusCode)
	w.Header().Set(contentTypeKey, jsonMime)
	_ = writeJSON(w, res)
}
