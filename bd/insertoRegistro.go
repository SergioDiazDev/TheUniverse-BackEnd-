package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/SergioDiazRuiz/TheUniverse/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertoRegistro Es la conexión final con la BD para insertar usuarios
func InsertoRegistro(u models.Usuario) (primitive.ObjectID, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("theuniverse")
	col := db.Collection("usuarios")

	u.Pass, _ = EncriptarPass(u.Pass)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return primitive.ObjectID{}, false, err
	}
	fmt.Println(u.ID)
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	u.ID = ObjID
	fmt.Println(ObjID)
	fmt.Println(u.ID)

	return ObjID, true, nil
}
