package jwt

import (
	"time"

	"github.com/SergioDiazRuiz/TheUniverse/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//GeneroJWT genera el encriptado con JWT
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("TheUniverse.es")

	payload := jwt.MapClaims{
		"correo":          t.Correo,
		"nombre":          t.Nombre,
		"fechaNacimiento": t.FechaNacimiento,
		"_id":             t.ID.Hex(),
		"exp":             time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
