package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/SergioDiazRuiz/TheUniverse/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ConsultoRelacion consulta la relacion entre dos usaurios
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := MongoCN.Database("theuniverse")
	col := db.Collection("relacion")

	var resultado models.Relacion

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"UsuarioRelacionID": t.UsuarioRelacionID,
	}

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
