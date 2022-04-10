package bd

import (
	"context"
	"time"

	"github.com/SergioDiazRuiz/TheUniverse/models"
)

func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoCN.Database("theuniverse")
	col := bd.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}
	return true, nil
}
