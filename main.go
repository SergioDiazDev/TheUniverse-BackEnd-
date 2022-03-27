package main

import (
	"log"

	"github.com/SergioDiazRuiz/TheUniverse/bd"
	"github.com/SergioDiazRuiz/TheUniverse/handlers"
)

func main() {
	if !bd.CheckConnection() {
		log.Fatal("Sin conxión a la BD")
		return
	}
	handlers.Gestor()
}
