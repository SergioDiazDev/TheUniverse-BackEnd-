package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/SergioDiazRuiz/TheUniverse/bd"
	"github.com/SergioDiazRuiz/TheUniverse/models"
)

//SubirAvatar subir Avatar al servidor
func SubirAvatar(w http.ResponseWriter, r *http.Request) {

	var id string = IDUsuario.Hex()
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + id + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	//defer f.Close()
	if err != nil {
		http.Error(w, "[Error] al subir avatar: "+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error Copiando avartar "+err.Error(), http.StatusBadRequest)
		return
	}
	var usuario models.Usuario
	var status bool

	usuario.Avatar = id + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)

	if err != nil || status == false {
		http.Error(w, "Error al grabar el avatar en la BD  "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
