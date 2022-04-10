package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SergioDiazRuiz/TheUniverse/bd"
	"github.com/SergioDiazRuiz/TheUniverse/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//consultaRelacion chequea si hay relacion entre 2 user
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion

	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID, _ = primitive.ObjectIDFromHex(ID)

	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(t)

	if err != nil || status == false {
		fmt.Println("Error" + err.Error())
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)
}
