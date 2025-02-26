package routes

import (
	"net/http"

	"github.com/faizkhan-06/go-auth/handlers"
	"github.com/faizkhan-06/go-auth/middlewares"
)

func RegisterRoutes() *http.ServeMux {
	router := http.NewServeMux();

	router.HandleFunc("POST /register",handlers.Register)
	router.HandleFunc("POST /login",handlers.Login)

	router.HandleFunc("GET /", middlewares.AuthMiddleware(handlers.Home))

	return router
}