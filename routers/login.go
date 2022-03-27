package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SergioDiazRuiz/TheUniverse/bd"
	"github.com/SergioDiazRuiz/TheUniverse/jwt"
	"github.com/SergioDiazRuiz/TheUniverse/models"
)

//Login realiza el login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o Contraseña incorrectas"+err.Error(), 400)
		return
	}
	if len(t.Correo) == 0 {
		http.Error(w, "El correo del usuario es requerido", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Correo, t.Pass)

	if !existe {
		http.Error(w, "Usuario y/o Contraseña incorrectas", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el Token "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//grabar cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
