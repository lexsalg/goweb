package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lexsalg/goweb/rest/models"
	"net/http"
	"strconv"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "LISTADO DE USUARIOS")
}

func Get(w http.ResponseWriter, r *http.Request) {

	//xml
	//user := models.User{Id: 1, Username: "lexsalg", Password: "123"}
	//w.Header().Set("Content-Type", "text/xml")
	//x, _ := xml.Marshal(&user)
	//_, _ = fmt.Fprintf(w, string(x))

	params := mux.Vars(r)
	userId, _ := strconv.Atoi(params["id"])
	res := models.DefaultResponse(w)
	user, err := models.GetUser(userId)
	if err != nil {
		res.NotFound()
	} else {
		res.Data = user
	}
	res.Send()
}

func Create(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "crear USUARIO")
}

func Update(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "update USUARIO")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "delete USUARIO")
}
