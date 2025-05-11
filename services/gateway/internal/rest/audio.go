package rest

import (
	"net/http"

	"github.com/what-da-flac/wtf/common/commands"
	"github.com/what-da-flac/wtf/common/restful"
	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

func (x *Server) PatchV1AudioFilesId(w http.ResponseWriter, r *http.Request, id string) {
	// read audio file from db
	audioFile, err := x.repository.SelectAudioFile(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	payload, err := restful.ReadRequest[golang.PatchV1AudioFilesIdJSONRequestBody](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	filename := "TODO"
	//  update file tags
	if err = commands.CmdFFMpegSetTags(x.identifier, filename, payload.Map()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// update db record
	if err = x.repository.UpdateAudioFile(id, payload.DBMap()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// read audio file from db
	if audioFile, err = x.repository.SelectAudioFile(id); err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// return audio_file record to caller
	restful.WriteJSONResponse(w, audioFile)
}
