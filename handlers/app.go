package handlers

import (
	"github.com/lexsalg/goweb/utils"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "app/index", nil)
}
