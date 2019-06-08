package v1

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lexsalg/goweb/models"
	"github.com/pkg/errors"
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
	user := &models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	if err := user.Valid(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	_ = user.SetPassword(user.Password)
	if err := user.Save(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	models.SendData(w, user)
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
	if err := user.Valid(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	user.Username = u.Username
	user.Email = u.Email
	_ = user.SetPassword(u.Password)

	if err := user.Save(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	models.SendData(w, user)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if user, err := getByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		_ = user.Delete()
		models.SendNoContent(w)
	}
}

func getByRequest(r *http.Request) (*models.User, error) {
	params := mux.Vars(r)
	userId, _ := strconv.Atoi(params["id"])
	user := models.GetUserById(userId)
	if user.Id == 0 {
		return user, errors.New("El usuario no existe.")
	}
	return user, nil

}
