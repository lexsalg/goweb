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
	r.Message = "Recurso No Encontrado."
}

func (r *Response) Send() {
	r.writer.Header().Set("Content-Type", r.contentType)
	r.writer.WriteHeader(r.Status)

	j, _ := json.Marshal(&r)
	_, _ = fmt.Fprintln(r.writer, string(j))
}

func (r *Response) UnprocessableEntity() {
	r.Status = http.StatusUnprocessableEntity
	r.Message = "Entidad No Reconocida."
}

func (r *Response) NoContent() {
	r.Status = http.StatusNoContent
	r.Message = "Sin Contenido."
}

/*---------------------------------------------------------------------*/
func SendData(w http.ResponseWriter, data interface{}) {
	response := DefaultResponse(w)
	response.Data = data
	response.Send()
}

func SendNotFound(w http.ResponseWriter) {
	response := DefaultResponse(w)
	response.NotFound()
	response.Send()
}

func SendUnprocessableEntity(w http.ResponseWriter) {
	response := DefaultResponse(w)
	response.UnprocessableEntity()
	response.Send()
}

func SendNoContent(w http.ResponseWriter) {
	response := DefaultResponse(w)
	response.NoContent()
	response.Send()
}
