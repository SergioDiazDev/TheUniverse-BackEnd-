package routers

import (
	"errors"
	"strings"

	"github.com/SergioDiazRuiz/TheUniverse/bd"
	"github.com/SergioDiazRuiz/TheUniverse/models"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Correo valor de correo usado en todos los EndPoints
var Correo string

//IDUsuario es el ID devuelto del modelo. que se usara en todos los EndPoints
var IDUsuario primitive.ObjectID

//ProcesoToken proceso token para extraer sus valores
func ProcesoToken(tk string) (*models.Claim, bool, primitive.ObjectID, error) {
	miClave := []byte("TheUniverse.es")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, primitive.NilObjectID, errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.CheckYaExisteUsuario(claims.Correo)

		if encontrado {
			IDUsuario = claims.ID
			Correo = claims.Correo
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, primitive.NilObjectID, errors.New("token invalido")
	}

	return claims, false, primitive.NilObjectID, err
}
