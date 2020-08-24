package main

import (
	"log"

	"github.com/jeffqev/tkey-api/database"
	"github.com/jeffqev/tkey-api/router"
)

func main() {

	if database.PingConnection() == 0 {
		log.Fatal("Sin Conexion")
		return
	}

	router.Serverup()

}
