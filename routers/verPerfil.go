package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SergioDiazRuiz/TheUniverse/bd"
)

//VerPerfil permite extraer los valores del Perfil
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	fmt.Println(ID)
	if len(ID) < 1 {

		http.Error(w, "Debe enviar el paramentro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Ocurrio un erro al buscar el registro"+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(perfil)

}
