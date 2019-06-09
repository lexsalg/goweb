package models

import "github.com/pkg/errors"

type ValidationError error

var (
	errUsername      = ValidationError(errors.New("El Usuario no debe estar vacio."))
	errShortUsername = ValidationError(errors.New("El Usuario debe tener min 3 letras."))
	errLargeUsername = ValidationError(errors.New("El Usuario debe tener max 30 letras."))

	errEmail              = ValidationError(errors.New("Correo no valido."))
	errPasswordEncryption = ValidationError(errors.New("no es posible cifrar la contraseña"))
	errLogin              = ValidationError(errors.New("Usuario o password no valido"))
)
