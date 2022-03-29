package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/SergioDiazRuiz/TheUniverse/middlew"
	"github.com/SergioDiazRuiz/TheUniverse/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Gestor seteo el puerto, el handler y escucho al server
func Gestor() {

	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.CheckBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
