package bd

import (
	"context"
	"time"

	"github.com/SergioDiazRuiz/TheUniverse/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//insertoFoto se encarga de insertar fotos
func InsertoFoto(t models.GraboFoto) (primitive.ObjectID, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("theuniverse")
	col := db.Collection("fotos")

	registro := bson.M{
		"usserid": t.UserID,
		"foto":    t.Foto,
		"fecha":   t.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return primitive.NilObjectID, false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID, true, nil
}
