package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Usuario modelo usuario de la BD
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Nombre          string             `bson:"nombre",  json:"nombre",omitempty`
	FechaNacimineto time.Time          `bson:"fechaNacimineto" json:"fechaNacimineto",omitempty`
	FechaRegistro   time.Time          `bson:"fechaRegistro" json:"fechaRegistro",omitempty`
	Correo          string             `bson:"correo" json:"correo"`
	Pass            string             `bson:"contraseña" json:"contraseña",omitempty`
	Avatar          string             `bson:"avatar" json:"avatar",omitempty`
	Banner          string             `bson:"banner" json:"banner",omitempty`
	Biografia       string             `bson:"biografia" json:"biografia",omitempty`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion",omitempty`
	SitioWeb        string             `bson:"sitioWeb" json:"sitioWeb",omitempty`
}
