package rest

import (
	"net/http"

	"github.com/what-da-flac/wtf/go-common/commands"

	"github.com/what-da-flac/wtf/openapi/domains"
)

func (x *Server) UploadAudioFile(w http.ResponseWriter, r *http.Request) {
	const fileFieldName = "file"
	err := r.ParseMultipartForm(100 << 20) // limit is 100 MB
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

	f := &domains.File{
		Id:          x.identifier.UUIDv4(),
		Filename:    filename,
		Created:     x.timer.Now(),
		Length:      size,
		ContentType: mimeType,
		Status:      domains.FileCreated.String(),
	}

	// send file content to storage implementation
	newFilename, err := x.fileStorage.Save(f, file)
	if err != nil {
		http.Error(w, "unable to save file", http.StatusInternalServerError)
		return
	}

	// extract mediainfo
	infoReader, err := commands.CmdMediaInfo(newFilename)
	if err != nil {
		http.Error(w, "unable to save file", http.StatusInternalServerError)
		return
	}
	info := domains.NewMediaInfo(infoReader)

	// convert mediainfo to audio data
	audio := domains.NewAudio(info)

	_ = audio

	if err = x.repository.InsertFile(f); err != nil {
		http.Error(w, "unable to save file metadata", http.StatusInternalServerError)
		return
	}

	// Respond with JSON (or store/save as needed)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
