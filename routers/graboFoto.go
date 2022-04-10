package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SergioDiazRuiz/TheUniverse/bd"
	"github.com/SergioDiazRuiz/TheUniverse/models"
)

//GraboFoto permite registrar una foto en la BD
func GraboFoto(w http.ResponseWriter, r *http.Request) {
	var foto models.Foto

	err := json.NewDecoder(r.Body).Decode(&foto)

	registro := models.GraboFoto{
		UserID: IDUsuario,
		Foto:   foto.Foto,
		Fecha:  time.Now(),
	}

	_, status, err := bd.InsertoFoto(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se a logrado insertar la foto ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
