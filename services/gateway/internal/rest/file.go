package rest

import (
	"net/http"

	"github.com/what-da-flac/wtf/openapi/domains"
)

func (x *Server) UploadAudioFile(w http.ResponseWriter, r *http.Request) {
	const fileFieldName = "file"
	// Parse up to 50 MB of incoming data (adjust if needed)
	err := r.ParseMultipartForm(50 << 20) // 50 MB
	if err != nil {
		http.Error(w, "unable to parse form", http.StatusBadRequest)
		return
	}

	// Get file from the form field named "file"
	file, fileHeader, err := r.FormFile(fileFieldName)
	if err != nil {
		http.Error(w, "file not found in request", http.StatusBadRequest)
		return
	}
	defer func() { _ = file.Close() }()

	// read file metadata
	filename := fileHeader.Filename
	size := fileHeader.Size
	mimeType := fileHeader.Header.Get("Content-Type")

	// TODO: send file content to storage implementation

	f := &domains.File{
		Id:          x.identifier.UUIDv4(),
		Filename:    filename,
		Created:     x.timer.Now(),
		Length:      size,
		ContentType: mimeType,
		Status:      domains.FileCreated.String(),
	}
	if err = x.repository.InsertFile(f); err != nil {
		http.Error(w, "unable to save file metadata", http.StatusInternalServerError)
		return
	}

	// Respond with JSON (or store/save as needed)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
