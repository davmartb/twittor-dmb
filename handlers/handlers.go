package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/davmartb/twittor-dmb/middlew"
	"github.com/davmartb/twittor-dmb/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Funcion Manejadores*/
func Manejadores() {
	router := mux.NewRouter()
	log.Println("Configurando manejadores")
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Println("Manejadores configurados")

	log.Println("Arrancado Server...")
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
