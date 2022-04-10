package bd

import (
	"context"
	"errors"
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
	//No se puede pasar la contraseña sin encriptar v/
	if len(u.Pass) > 0 {
		PassEn, _ := EncriptarPass(u.Pass)
		registro["contraseña"] = PassEn
	}
	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}

	_, correoUsado, _ := CheckYaExisteUsuario(u.Correo)

	t, _ := BuscoPerfil(ID.Hex())

	//Cambio el todo menos el correo si es el mismo o esta ya registrado
	//Problema: No muestra el error de que el correo ya esta registrado
	var errr error
	if !correoUsado {
		if len(u.Correo) > 0 {
			registro["correo"] = u.Correo
		}

	} else {
		if t.Correo == u.Correo {
			//Modifica correo pero muestra el error
			registro["correo"] = u.Correo
			errr = errors.New(" Este es tu correo actual")
		} else {
			errr = errors.New(" Este correo ya se encuentra registrado")
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

	return true, errr
}
