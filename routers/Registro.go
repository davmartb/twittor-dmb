package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/davmartb/twittor-dmb/bd"
	"github.com/davmartb/twittor-dmb/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Fatal("Error de formato de json de Registro, " + err.Error())
		http.Error(w, "Error en los datos recibido, "+err.Error(), http.StatusBadRequest)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "el mail del usuario es requerido", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una password de seis caracteres", http.StatusBadRequest)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario regitrado con ese mail", http.StatusBadRequest)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro del usuario", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro del usuario", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
