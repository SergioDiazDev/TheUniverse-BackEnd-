package bd

import (
	"context"
	"time"

	"github.com/SergioDiazRuiz/TheUniverse/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ModificoRegistro permite modificar el perfil de usuario
func ModificoRegistro(u models.Usuario, ID primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoCN.Database("theuniverse")
	col := bd.Collection("usuarios")

	registro := make(map[string]interface{})
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	registro["fechaNacimiento"] = u.FechaNacimiento
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Correo) > 0 {
		registro["correo"] = u.Correo
	}
	//No se puede pasar la contraseña sin encriptar
	if len(u.Pass) > 0 {
		registro["contraseña"] = u.Pass
	}
	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}
	//Error de conversion intervace string
	correo := registro["correo"]
	correoString := correo.(string)
	_, correoUsado, _ := CheckYaExisteUsuario(string(correoString))

	if correoUsado {
		if u.Correo == correoString {
			registro["correo"] = u.Correo

		} else {
			return false, nil
		}
	}

	updateString := bson.M{
		"$set": registro,
	}

	filtro := bson.M{"_id": bson.M{"$eq": ID}}

	_, err := col.UpdateOne(ctx, filtro, updateString)

	if err != nil {
		return false, err
	}
	return true, nil
}
