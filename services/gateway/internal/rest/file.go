package rest

import (
	"fmt"
	"net/http"

	"github.com/what-da-flac/wtf/openapi/domains"
	"github.com/what-da-flac/wtf/openapi/gen/golang"
	"golang.org/x/net/context"
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
		x.logger.Errorf("unable to get file from form: %v", err)
		http.Error(w, "file not found in request", http.StatusBadRequest)
		return
	}
	defer func() { _ = file.Close() }()

	// read file metadata
	contentType := domains.NewContentType(fileHeader.Header.Get("Content-Type"))
	if contentType == golang.ContentTypeInvalid {
		err = fmt.Errorf("unsupported media type %q", fileHeader.Header.Get("Content-Type"))
		x.logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// send file content to storage implementation
	filename := x.identifier.UUIDv4()
	if err = x.tempPathFinder.Save(file, filename); err != nil {
		x.logger.Errorf("unable to save file: %v", err)
		http.Error(w, "unable to save file", http.StatusInternalServerError)
		return
	}
	ctx := context.Background()
	if _, err = x.mediaInfoPublisher.PublishMessage(ctx, golang.MediaInfoInput{
		ContentType:            contentType,
		Filename:               filename,
		OriginalFilename:       fileHeader.Filename,
		SrcPathName:            x.tempPathFinder.Path(),
		DstPathName:            x.storePathFinder.Path(),
		MinBitrate:             192 * 1000, // TODO: this should come from somewhere else
		ConvertedBitRate:       320 * 1000, // TODO: this should come from somewhere else
		DestinationContentType: golang.ContentTypeAudioxM4a,
	}); err != nil {
		x.logger.Errorf("unable to publish audio file: %v", err)
		http.Error(w, "unable to publish audio file", http.StatusInternalServerError)
		return
	}
}
