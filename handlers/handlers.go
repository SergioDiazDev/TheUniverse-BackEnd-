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
	router.HandleFunc("/modificarPerfil", middlew.CheckBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/foto", middlew.CheckBD(middlew.ValidoJWT(routers.GraboFoto))).Methods("POST")
	router.HandleFunc("/subirAvatar", middlew.CheckBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.CheckBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/altaRelacion", middlew.CheckBD(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.CheckBD(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.CheckBD(middlew.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")
	router.HandleFunc("/listaUsuarios", middlew.CheckBD(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
