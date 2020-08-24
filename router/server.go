package router

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

/*Serverup Levanta el servidor junto a los handler que se encuentren definidos, en el puerto por defecto 4000 o segun la Variable de entorno */
func Serverup() {
	route := mux.NewRouter()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Cargando Variable de entorno")
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "4000"
	}

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome Server UP")
	})

	UserHandlers(route)

	handler := cors.AllowAll().Handler(route)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
