package middlew

import (
	"net/http"

	"github.com/SergioDiazRuiz/TheUniverse/bd"
)

//CheckBD Se encarga de comprobar la conexión
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !bd.CheckConnection() {
			http.Error(w, "Conexión perdida con BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
