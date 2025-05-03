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
	if contentType == golang.MediaInfoInputContentTypeInvalid {
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
		ContentType:      contentType,
		Filename:         filename,
		OriginalFilename: fileHeader.Filename,
		PathName:         x.tempPathFinder.Path(),
		MinBitrate:       320 * 1000, // TODO: this should come from somewhere else
	}); err != nil {
		x.logger.Errorf("unable to publish audio file: %v", err)
		http.Error(w, "unable to publish audio file", http.StatusInternalServerError)
		return
	}
	//return
	//
	//f := &golang.File{
	//	Id:          x.identifier.UUIDv4(),
	//	Filename:    filename,
	//	Created:     x.timer.Now(),
	//	Length:      int(size),
	//	ContentType: string(contentType),
	//	Status:      domains.FileCreated.String(),
	//}
	//
	//convert to aac audio format
	//const m4aExt = ".m4a"
	//ext := filepath.Ext(filename)
	//dstFilename := srcFilename
	//dir := filepath.Dir(srcFilename)
	//if ext != m4aExt {
	//	base := filepath.Base(srcFilename)
	//	dstFilename = strings.TrimSuffix(base, ext)
	//	dstFilename = filepath.Join(dir, dstFilename+m4aExt)
	//	if err = commands.CmdFFMpegAudio(srcFilename, dstFilename); err != nil {
	//		x.logger.Errorf("unable to convert to audio file: %s", err)
	//		http.Error(w, "unable to convert to audio file", http.StatusInternalServerError)
	//		return
	//	}
	//	file size needs to be calculated after conversion
	//if si, err := os.Stat(dstFilename); err != nil {
	//	x.logger.Errorf("unable to stat audio file: %s", err)
	//	http.Error(w, "file does not exist", http.StatusConflict)
	//} else {
	//	f.Length = int(si.Size())
	//}
	//}
	//f.Filename = filepath.Base(dstFilename)
	//
	//extract mediainfo
	//infoReader, err := commands.CmdMediaInfo(dstFilename)
	//if err != nil {
	//	x.logger.Errorf("unable to get media info: %s", err)
	//	http.Error(w, "unable to save file", http.StatusInternalServerError)
	//	return
	//}
	//info, err := domains.NewMediaInfo(infoReader)
	//if err != nil {
	//	http.Error(w, "unable to get media info file", http.StatusInternalServerError)
	//	return
	//}
	//
	//convert mediainfo to audio data
	//audio := domains.NewAudio(info)
	//
	//save audio file to db
	//audioFile := domains.NewAudioFile(&audio, f)
	//
	//if err = x.repository.InsertAudioFile(&audioFile); err != nil {
	//	x.logger.Errorf("unable to save audio file: %s", err)
	//	http.Error(w, "unable to save file metadata", http.StatusInternalServerError)
	//	return
	//}
	//Respond with JSON (or store/save as needed)
	//http_helpers.WriteJSON(w, http.StatusOK, audio)
}
