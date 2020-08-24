package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoConn Coloca a disposicion global la conexion*/
var MongoConn = ConnectDB()

/*ConnectDB Funcion para conectar a la base de datos*/
func ConnectDB() *mongo.Client {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Cargando Variable de entorno")

	}

	mongoURL := os.Getenv("MONGODB")

	clientOption := options.Client().ApplyURI(mongoURL)
	// clientOption := options.Client().ApplyURI(os.Getenv("MONGODB"))

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}

	log.Println("Conexion exitosa")

	return client

}

/*PingConnection revisa que la conexion continue a la base de datos por medio de un ping, esto debido a que la conexion quedara de manera global en la variable MongoConn */
func PingConnection() int {
	err := MongoConn.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return 0
	}

	return 1
}
