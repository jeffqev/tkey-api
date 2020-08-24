package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*UserHandlers Contiene las rutas de los usuarios del API */
func UserHandlers(route *mux.Router) {
	route.HandleFunc("/holaa", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "desde users")
	})

}
