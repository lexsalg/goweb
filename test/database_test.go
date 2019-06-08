package test

import (
	"testing"

	"github.com/lexsalg/goweb/models"
)

func TestConecction(t *testing.T) {

	if conn := models.GetConnection(); conn == nil {
		t.Error("no es posible realizar la conexion", nil)
	}

}
