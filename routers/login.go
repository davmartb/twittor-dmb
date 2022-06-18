package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/davmartb/twittor-dmb/bd"
	"github.com/davmartb/twittor-dmb/jwt"
	"github.com/davmartb/twittor-dmb/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o Contraseña inválidos"+err.Error(), http.StatusBadRequest)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "el mail del usuario es requerido", http.StatusBadRequest)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		http.Error(w, "El usuario no existe", http.StatusNotFound)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar ek Token correspondiente"+err.Error(), http.StatusBadRequest)
		return
	}
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
