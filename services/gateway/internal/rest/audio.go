package rest

import (
	"encoding/json"
	"net/http"

	"github.com/what-da-flac/wtf/common/commands"

	"github.com/what-da-flac/wtf/common/restful"
)

func (x *Server) PatchV1AudioFilesId(w http.ResponseWriter, r *http.Request, id string) {
	// read audio file from db
	audioFile, err := x.repository.SelectAudioFile(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// extract map of fields from body
	var m map[string]any
	if err = json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// remove fields we don't want to update
	for _, f := range []string{
		"id", "compression_mode", "content_type",
		"created", "duration_seconds", "file_extension",
		"filename", "format", "genre",
		"sampling_rate",
	} {
		delete(m, f)
	}

	//  TODO: update file tags
	commands.CmdFFMpegSetTags(x.identifier, filename, m)
	// TODO: update db record
	// select db record
	if audioFile, err = x.repository.SelectAudioFile(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return audio_file record to caller
	restful.WriteJSONResponse(w, audioFile)
}
