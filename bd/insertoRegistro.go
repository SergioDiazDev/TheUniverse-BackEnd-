package bd

import (
	"context"
	"time"

	"github.com/SergioDiazRuiz/TheUniverse/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertoRegistro Es la conexión final con la BD para insertar usuarios
func InsertoRegistro(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("theuniverse")
	col := db.Collection("usuarios")

	u.Pass, _ = EncriptarPass(u.Pass)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
