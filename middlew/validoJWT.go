package middlew

import "net/http"

func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
