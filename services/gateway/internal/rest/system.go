package rest

import (
	"encoding/json"
	"net/http"

	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

func (x *Server) GetV1Healthz(w http.ResponseWriter, r *http.Request) {
	res := &golang.Health{
		Ok:      true,
		Version: "dev",
	}
	data, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data = []byte(err.Error())
		return
	}
	_, _ = w.Write(data)
}

func (x *Server) GetV1Container(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
