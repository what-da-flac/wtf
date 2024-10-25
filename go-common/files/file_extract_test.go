package files

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestPayload struct {
	Key string `json:"key"`
}

func TestExtractFileFromRequest(t *testing.T) {
	// thanks chatgpt :D
	t.Run("Successful extraction", func(t *testing.T) {
		// Create a buffer to write our multipart form data
		var b bytes.Buffer
		w := multipart.NewWriter(&b)

		// Create a form file field
		fileWriter, err := w.CreateFormFile("fileKey", "testfile.txt")
		assert.NoError(t, err)

		// Write some data to the file field
		_, err = io.WriteString(fileWriter, "This is a test file content")
		assert.NoError(t, err)

		// Create a form field for the JSON payload
		payload := TestPayload{Key: "value"}
		payloadBytes, err := json.Marshal(payload)
		assert.NoError(t, err)
		err = w.WriteField("jsonKey", string(payloadBytes))
		assert.NoError(t, err)

		// Close the writer
		err = w.Close()
		assert.NoError(t, err)

		// Create a new HTTP request with the multipart form data
		req := httptest.NewRequest(http.MethodPost, "http://example.com/upload", &b)
		req.Header.Set("Content-Type", w.FormDataContentType())

		// Call the function under test
		var extractedPayload TestPayload
		filename, fileSize, file, err := ExtractFileFromRequest(req, "fileKey", "jsonKey", &extractedPayload)
		assert.NoError(t, err)

		// Verify the results
		assert.Equal(t, "testfile.txt", filename)
		assert.Equal(t, int64(len("This is a test file content")), fileSize)
		assert.Equal(t, payload, extractedPayload)

		// Verify file content
		fileContent, err := io.ReadAll(file)
		assert.NoError(t, err)
		assert.Equal(t, "This is a test file content", string(fileContent))
	})

	t.Run("Missing file part", func(t *testing.T) {
		// Create a buffer to write our multipart form data
		var b bytes.Buffer
		w := multipart.NewWriter(&b)

		// Create a form field for the JSON payload
		payload := TestPayload{Key: "value"}
		payloadBytes, err := json.Marshal(payload)
		assert.NoError(t, err)
		err = w.WriteField("jsonKey", string(payloadBytes))
		assert.NoError(t, err)

		// Close the writer
		err = w.Close()
		assert.NoError(t, err)

		// Create a new HTTP request with the multipart form data
		req := httptest.NewRequest(http.MethodPost, "http://example.com/upload", &b)
		req.Header.Set("Content-Type", w.FormDataContentType())

		// Call the function under test
		var extractedPayload TestPayload
		_, _, _, err = ExtractFileFromRequest(req, "fileKey", "jsonKey", &extractedPayload)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "missing file part")
	})

	t.Run("Invalid JSON payload", func(t *testing.T) {
		// Create a buffer to write our multipart form data
		var b bytes.Buffer
		w := multipart.NewWriter(&b)

		// Create a form file field
		fileWriter, err := w.CreateFormFile("fileKey", "testfile.txt")
		assert.NoError(t, err)

		// Write some data to the file field
		_, err = io.WriteString(fileWriter, "This is a test file content")
		assert.NoError(t, err)

		// Create an invalid JSON payload
		err = w.WriteField("jsonKey", "invalid json")
		assert.NoError(t, err)

		// Close the writer
		err = w.Close()
		assert.NoError(t, err)

		// Create a new HTTP request with the multipart form data
		req := httptest.NewRequest(http.MethodPost, "http://example.com/upload", &b)
		req.Header.Set("Content-Type", w.FormDataContentType())

		// Call the function under test
		var extractedPayload TestPayload
		_, _, _, err = ExtractFileFromRequest(req, "fileKey", "jsonKey", &extractedPayload)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid character")
	})
}
