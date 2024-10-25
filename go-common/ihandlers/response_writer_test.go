package ihandlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewResponseWriter(t *testing.T) {
	rw := httptest.NewRecorder()
	w := NewResponseWriter(rw)
	if w.statusCode != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, w.statusCode)
	}
	if w.err != nil {
		t.Errorf("expected nil error, got %v", w.err)
	}
}

func TestWriteHeader(t *testing.T) {
	rw := httptest.NewRecorder()
	w := NewResponseWriter(rw)

	// Test setting a custom status code
	w.WriteHeader(http.StatusNotFound)
	if w.statusCode != http.StatusNotFound {
		t.Errorf("expected status code %v, got %v", http.StatusNotFound, w.statusCode)
	}

	// Test that status code can only be assigned once
	w.WriteHeader(http.StatusInternalServerError)
	if w.statusCode != http.StatusNotFound {
		t.Errorf("expected status code %v, got %v", http.StatusNotFound, w.statusCode)
	}
}

func TestWriteHeaderDefaultStatusCode(t *testing.T) {
	rw := httptest.NewRecorder()
	w := NewResponseWriter(rw)

	// Test setting the default status code (should be ignored)
	w.WriteHeader(http.StatusOK)
	if w.statusCode != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, w.statusCode)
	}
}

func TestWrite(t *testing.T) {
	rw := httptest.NewRecorder()
	w := NewResponseWriter(rw)

	// Test writing a response body
	data := []byte("hello, world")
	n, err := w.Write(data)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if n != len(data) {
		t.Errorf("expected %v bytes written, got %v", len(data), n)
	}
	if rw.Body.String() != string(data) {
		t.Errorf("expected body %v, got %v", string(data), rw.Body.String())
	}

	// Test writing with an existing error
	w.err = http.ErrBodyNotAllowed
	n, err = w.Write(data)
	if err != http.ErrBodyNotAllowed {
		t.Errorf("expected error %v, got %v", http.ErrBodyNotAllowed, err)
	}
	if n != 0 {
		t.Errorf("expected 0 bytes written, got %v", n)
	}
}

func TestIsError(t *testing.T) {
	rw := httptest.NewRecorder()
	w := NewResponseWriter(rw)

	// Test that default status code is not an error
	if w.IsError() {
		t.Errorf("expected no error, got error")
	}

	// Test that status code >= 400 is an error
	w.WriteHeader(http.StatusInternalServerError)
	if !w.IsError() {
		t.Errorf("expected error, got no error")
	}
}

func TestStatusCode(t *testing.T) {
	rw := httptest.NewRecorder()
	w := NewResponseWriter(rw)

	// Test getting the status code
	if w.StatusCode() != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, w.StatusCode())
	}

	// Test getting a custom status code
	w.WriteHeader(http.StatusBadRequest)
	if w.StatusCode() != http.StatusBadRequest {
		t.Errorf("expected status code %v, got %v", http.StatusBadRequest, w.StatusCode())
	}
}
