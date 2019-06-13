package handlers

import (
	"net/http"

	"github.com/lexsalg/goweb/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "app/index", nil)
}

func App1(w http.ResponseWriter, r *http.Request) {
	utils.RenderApp(w, "correo-else")
}

func App2(w http.ResponseWriter, r *http.Request) {
	utils.RenderApp(w, "tramite")
}

func App3(w http.ResponseWriter, r *http.Request) {
	utils.RenderApp(w, "clima")
}
