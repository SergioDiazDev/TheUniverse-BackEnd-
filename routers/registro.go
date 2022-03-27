package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SergioDiazRuiz/TheUniverse/bd"
	"github.com/SergioDiazRuiz/TheUniverse/models"
)

//Registro es la funcion encargada de registrar a los usuarios
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Correo) == 0 {
		http.Error(w, "Email de usuario requerido", 400)
		return
	}
	if len(t.Pass) < 8 {
		http.Error(w, "La contraseña debe tener minimo 8 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.CheckYaExisteUsuario(t.Correo)

	if encontrado {
		http.Error(w, "Ya existe un usuario con este Email", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if status {
		http.Error(w, "No se a logrado insertar el registro de usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
