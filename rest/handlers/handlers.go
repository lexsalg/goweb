package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lexsalg/goweb/rest/models"
	"net/http"
	"strconv"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetUsers())
}

func Get(w http.ResponseWriter, r *http.Request) {

	if user, err := getByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		models.SendData(w, user)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		models.SendData(w, models.SaveUser(user))
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	user, err := getByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}

	u := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&u); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	user = models.UpdateUser(user, u.Username, u.Password)
	models.SendData(w, user)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if user, err := getByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		models.DeleteUser(user.Id)
		models.SendNoContent(w)
	}
}

func getByRequest(r *http.Request) (models.User, error) {
	params := mux.Vars(r)
	userId, _ := strconv.Atoi(params["id"])
	if user, err := models.GetUser(userId); err != nil {
		return user, err
	} else {
		return user, nil
	}
}
