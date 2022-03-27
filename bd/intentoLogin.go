package bd

import (
	"github.com/SergioDiazRuiz/TheUniverse/models"
	"golang.org/x/crypto/bcrypt"
)

//IntentoLogin realiza el check de login en la BD
func IntentoLogin(email string, pass string) (models.Usuario, bool) {
	usu, encontrado, _ := CheckYaExisteUsuario(email)
	if !encontrado {
		return usu, false
	}

	passBytes := []byte(pass)
	passBD := []byte(usu.Pass)

	err := bcrypt.CompareHashAndPassword(passBD, passBytes)

	if err != nil {
		return usu, false
	}
	return usu, true
}
