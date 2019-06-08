package test

import (
	"testing"

	"github.com/lexsalg/goweb/rest/models"
)

func TestConecction(t *testing.T) {

	if conn := models.GetConnection(); conn == nil {
		t.Error("no es posible realizar la conexion", nil)
	}

}
