package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GraboFoto es la estructura que tendran nuestras fotos
type GraboFoto struct {
	UserID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Foto   string             `bson:"foto" json:userid:"foto,omitempty"`
	Fecha  time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
