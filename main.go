package main

import (
	"fmt"
	"log"

	"github.com/davmartb/twittor-dmb/bd"
	"github.com/davmartb/twittor-dmb/handlers"
)

func main() {
	fmt.Println("Esto es una prueba")
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BBDD")
	}
	handlers.Manejadores()
}
