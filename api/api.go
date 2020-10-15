package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	allowedHeaders = []string{"X-Requested-With", "Content-Type", "Authorization"}
	allowedMethods = []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}
	allowedOrigins = []string{"*"}
)

// CreateRoutes sets up routes on a gorilla mux router
func CreateRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")

	return router
}

// CreateHTTPHandler sets up CORS and returns the http handler
func CreateHTTPHandler(router *mux.Router) http.Handler {
	return handlers.CORS(handlers.AllowedHeaders(allowedHeaders), handlers.AllowedMethods(allowedMethods), handlers.AllowedOrigins(allowedOrigins))(router)
}

func home(rw http.ResponseWriter, req *http.Request) {
	json.NewEncoder(rw).Encode(`{"response": "Hello world"}`)
}
