package rest

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/what-da-flac/wtf/openapi/gen/golang"

	"github.com/what-da-flac/wtf/go-common/commands"
	"github.com/what-da-flac/wtf/go-common/http_helpers"
	"github.com/what-da-flac/wtf/openapi/domains"
)

func (x *Server) UploadAudioFile(w http.ResponseWriter, r *http.Request) {
	const fileFieldName = "file"
	err := r.ParseMultipartForm(500 << 20) // limit is 500 MB
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

	f := &golang.File{
		Id:          x.identifier.UUIDv4(),
		Filename:    filename,
		Created:     x.timer.Now(),
		Length:      int(size),
		ContentType: mimeType,
		Status:      domains.FileCreated.String(),
	}

	// send file content to storage implementation
	srcFilename, err := x.fileStorage.Save(f, file)
	if err != nil {
		http.Error(w, "unable to save file", http.StatusInternalServerError)
		return
	}

	// convert to aac audio format
	const m4aExt = ".m4a"
	dstFilename := srcFilename
	dir := filepath.Dir(srcFilename)
	if ext := filepath.Ext(srcFilename); ext != m4aExt {
		base := filepath.Base(srcFilename)
		dstFilename = strings.TrimSuffix(base, ext)
		dstFilename = filepath.Join(dir, dstFilename+m4aExt)
		if err = commands.CmdFFMpegAudio(srcFilename, dstFilename); err != nil {
			x.logger.Errorf("unable to convert to audio file: %s", err)
			http.Error(w, "unable to convert to audio file", http.StatusInternalServerError)
			return
		}
		// file size needs to be calculated after conversion
		if si, err := os.Stat(dstFilename); err != nil {
			x.logger.Errorf("unable to stat audio file: %s", err)
			http.Error(w, "file does not exist", http.StatusConflict)
		} else {
			f.Length = int(si.Size())
		}
	}
	f.Filename = filepath.Base(dstFilename)

	// extract mediainfo
	infoReader, err := commands.CmdMediaInfo(dstFilename)
	if err != nil {
		http.Error(w, "unable to save file", http.StatusInternalServerError)
		return
	}
	info, err := domains.NewMediaInfo(infoReader)
	if err != nil {
		http.Error(w, "unable to get media info file", http.StatusInternalServerError)
		return
	}

	// convert mediainfo to audio data
	audio := domains.NewAudio(info)

	// save audio file to db
	audioFile := domains.NewAudioFile(&audio, f)

	if err = x.repository.InsertAudioFile(&audioFile); err != nil {
		http.Error(w, "unable to save file metadata", http.StatusInternalServerError)
		return
	}
	// Respond with JSON (or store/save as needed)
	http_helpers.WriteJSON(w, http.StatusOK, audio)
}
