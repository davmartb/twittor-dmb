package middlew

import (
	"net/http"

	"github.com/davmartb/twittor-dmb/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la Base de Datos", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
