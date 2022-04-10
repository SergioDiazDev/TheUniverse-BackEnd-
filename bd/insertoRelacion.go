package bd

import (
	"context"
	"time"

	"github.com/SergioDiazRuiz/TheUniverse/models"
)

//InsertoRelacion graba la relacionen la BD
func IsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()
	db := MongoCN.Database("theuniverse")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)

	if err != nil {
		return false, err
	}
	return true, nil
}
