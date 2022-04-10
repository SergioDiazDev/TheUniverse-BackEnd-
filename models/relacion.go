package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Relacion modelo para grabar la relacion de un usuario con otro
type Relacion struct {
	Id                primitive.ObjectID `bson:"_id" json:"_id"`
	UsuarioID         primitive.ObjectID `bson:"usuarioid" json:"usuarioid"`
	UsuarioRelacionID primitive.ObjectID `bson:"UsuarioRelacionID" json:"UsuarioRelacionID"`
}
