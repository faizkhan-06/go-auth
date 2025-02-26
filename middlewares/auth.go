package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/faizkhan-06/go-auth/utils"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(utils.Response{
				Message: "You are not authorized",
				Status:  http.StatusUnauthorized,
			})
			return
		}

		tokenString := authHeader[len("Bearer "):]
		err := utils.VerifyJWTToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(utils.Response{
				Message: "Invalid token",
				Status:  http.StatusUnauthorized,
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}