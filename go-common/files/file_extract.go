package files

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/what-da-flac/wtf/go-common/exceptions"
)

// ExtractFileFromRequest reads from http.request and extracts from a multipart form,
// the file metadata and the file reader itself.
// Parameter fileKey should be "file".
// Parameter jsonKey should be "json".
// A callback function is received to pass the resulting file from the multipart request.
func ExtractFileFromRequest(r *http.Request, fileKey, jsonKey string, payload any) (filename string, fileSize int64, file multipart.File, err error) {
	// TODO: set this value from environment
	// 10 Mb
	var maxMemory int64 = 10 << 20
	if err = r.ParseMultipartForm(maxMemory); err != nil {
		err = exceptions.NewHTTPError(err).WithStatusCode(http.StatusBadRequest)
		return
	}
	// read file content from request
	formFile, ok := r.MultipartForm.File[fileKey]
	if !ok || len(formFile) == 0 {
		err = exceptions.NewHTTPError(fmt.Errorf("missing file part")).WithStatusCode(http.StatusBadRequest)
		return
	}
	headerFile := formFile[0]

	// read json content from request
	formJSON, ok := r.MultipartForm.Value[jsonKey]
	if !ok || len(formJSON) == 0 {
		err = exceptions.NewHTTPError(fmt.Errorf("missing json part")).WithStatusCode(http.StatusBadRequest)
		return
	}
	headerPayload := formJSON[0]
	if payload != nil {
		if err = json.Unmarshal([]byte(headerPayload), payload); err != nil {
			err = exceptions.NewHTTPError(err).WithStatusCode(http.StatusBadRequest)
			return
		}
	}
	// open uploaded file
	file, err = headerFile.Open()
	if err != nil {
		defer func() { _ = file.Close() }()
		err = exceptions.NewHTTPError(err).WithStatusCode(http.StatusBadRequest)
		return
	}
	filename = headerFile.Filename
	fileSize = headerFile.Size
	return
}
