package bd

import (
	"context"
	"time"

	"github.com/SergioDiazRuiz/TheUniverse/models"
	"go.mongodb.org/mongo-driver/bson"
)

//CheckYaExisteUsuario recibe un email de paramentro y mira si existe en la BD
//|true->Exite ya el Correo
func CheckYaExisteUsuario(correo string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("theuniverse")
	col := db.Collection("usuarios")

	condicion := bson.M{"correo": correo}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
