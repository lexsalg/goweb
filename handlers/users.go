package handlers

import (
	"net/http"

	"github.com/lexsalg/goweb/models"
	"github.com/lexsalg/goweb/utils"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	ctx := make(map[string]interface{})

	if r.Method == "POST" {
		u := r.FormValue("username")
		e := r.FormValue("email")
		p := r.FormValue("password")
		if user, err := models.CreateUser(u, p, e); err != nil {
			ctx["Error"] = err.Error()
		} else {
			utils.SetSession(user, w)
			http.Redirect(w, r, "/users/edit", http.StatusSeeOther)
		}
	}
	utils.RenderTemplate(w, "users/new", ctx)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := make(map[string]interface{})
	user := utils.GetUser(r)
	ctx["User"] = user

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	utils.RenderTemplate(w, "users/edit", ctx)
}

func Login(w http.ResponseWriter, r *http.Request) {
	ctx := make(map[string]interface{})

	if r.Method == "POST" {
		u := r.FormValue("username")
		p := r.FormValue("password")
		if user, err := models.Login(u, p); err != nil {
			ctx["Error"] = err.Error()
		} else {
			utils.SetSession(user, w)
			http.Redirect(w, r, "/users/edit", http.StatusSeeOther)
		}
	}
	utils.RenderTemplate(w, "users/login", ctx)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	utils.DeleteSession(w, r)

	utils.RenderTemplate(w, "users/logout", nil)

	// http.Redirect(w, r, "/users/login?exit", http.StatusSeeOther)
}
