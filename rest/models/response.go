package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	contentType string
	writer      http.ResponseWriter
}

func DefaultResponse(w http.ResponseWriter) Response {
	contentType := "application/json"
	return Response{Status: http.StatusOK, writer: w, contentType: contentType}
}

func (r *Response) NotFound() {
	r.Status = http.StatusNotFound
	r.Data = nil
	r.Message = "Recurso no Encontrado"
}

func (r *Response) Send() {
	r.writer.WriteHeader(r.Status)
	r.writer.Header().Set("Content-Type", r.contentType)

	j, _ := json.Marshal(&r)
	_, _ = fmt.Fprintln(r.writer, string(j))

}
