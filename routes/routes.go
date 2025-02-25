package routes

import (
	"net/http"

	"github.com/faizkhan-06/go-auth/handlers"
)

func RegisterRoutes() *http.ServeMux {
	router := http.NewServeMux();

	router.HandleFunc("POST /register",handlers.Register)
	router.HandleFunc("POST /login",handlers.Login)

	return router
}