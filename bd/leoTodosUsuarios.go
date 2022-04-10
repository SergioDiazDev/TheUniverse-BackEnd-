package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/SergioDiazRuiz/TheUniverse/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//LeoTodosUsuarios Lee todos los usuarios registrados en el sistema
func LeoTodosUsuarios(ID primitive.ObjectID, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("theuniverse")
	col := db.Collection("usuarios")

	var result []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}
	cur, err := col.Find(ctx, query, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}

	var encontrado, incluir bool

	for cur.Next(ctx) {
		var s models.Usuario
		err := cur.Decode(&s)

		if err != nil {
			fmt.Println(err.Error())
			return result, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID

		incluir = false

		encontrado, err = ConsultoRelacion(r)
		if tipo == "new" && encontrado == false {
			incluir = true
		}
		if tipo == "follow" && encontrado == true {
			incluir = true
		}

		//Esto evita que los usuarios se vean chuparse la pinga
		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir == true {
			s.Pass = ""
			s.Biografia = ""
			s.Correo = ""
			s.Banner = ""
			s.FechaNacimiento = time.Time{}
			s.FechaRegistro = time.Time{}
			s.SitioWeb = ""
			s.Ubicacion = ""

			result = append(result, &s)
		}
	}
	err = cur.Err()

	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}
	cur.Close(ctx)
	return result, true
}
