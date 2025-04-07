package rest

import (
	"fmt"
	"net/http"
)

func (x *Server) UploadAudioFile(w http.ResponseWriter, r *http.Request) {
	// Parse up to 50 MB of incoming data (adjust if needed)
	err := r.ParseMultipartForm(50 << 20) // 50 MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get file from the form field named "file"
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File not found in request", http.StatusBadRequest)
		return
	}
	defer func() { _ = file.Close() }()

	// Read file content (example: first 512 bytes for MIME sniffing)
	buf := make([]byte, 512)
	n, _ := file.Read(buf)

	// Reset the reader to start (because we read from it)
	if _, err = file.Seek(0, 0); err != nil {
		http.Error(w, "File not found in request", http.StatusBadRequest)
		return
	}

	// Sniff the MIME type
	mimeType := http.DetectContentType(buf[:n])

	// Optional: read full file into memory (only if needed)
	// fullData, err := io.ReadAll(file)
	// if err != nil {
	//     http.Error(w, "Failed to read file", http.StatusInternalServerError)
	//     return
	// }

	// Log or use file metadata
	fmt.Printf("Filename: %s\n", fileHeader.Filename)
	fmt.Printf("Size: %d bytes\n", fileHeader.Size)
	fmt.Printf("MIME type: %s\n", mimeType)

	// Respond with JSON (or store/save as needed)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, `{"filename":"%s","size":%d,"mime":"%s"}`,
		fileHeader.Filename, fileHeader.Size, mimeType)
}
