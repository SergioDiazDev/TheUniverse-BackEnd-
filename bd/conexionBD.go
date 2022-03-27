package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN Objeto de concexión a la BD
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://TheUniverse:Sergioforo1@cluster0.gptrm.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

//ConectarBD: Funcion encargada de conectar la BD
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión exitosa con la BD")
	return client
}

//CheckConnection: Funcion encargada de hacer ping a la BD
func CheckConnection() bool {

	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
