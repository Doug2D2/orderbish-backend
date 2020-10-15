package main

import (
	"log"
	"net/http"

	"github.com/Doug2D2/orderbish-backend/api"
	"github.com/Doug2D2/orderbish-backend/db"
)

func main() {
	db.Setup()
	router := api.CreateRoutes()
	handler := api.CreateHTTPHandler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
