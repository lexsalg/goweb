package test

import "testing"

func TestHolaMundo(t *testing.T) {
	str := "hola Mundo"
	if str != "hola Mundo" {
		t.Error("no es posible saludar al usuario", nil)
	}
}

func TestHolaMundo2(t *testing.T) {
	str := "hola Mundo2"
	if str != "hola Mundo2" {
		t.Error("no es posible saludar al usuario 2", nil)
	}
}
