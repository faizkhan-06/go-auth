package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/faizkhan-06/go-auth/config"
	"github.com/faizkhan-06/go-auth/routes"
)

func main() {
	config.ConnectDb();

	router := routes.RegisterRoutes()
	fmt.Println("server is running on port 4000")
	err := http.ListenAndServe(":4000", router)
	if err != nil {
		log.Fatal(err)
	}
}