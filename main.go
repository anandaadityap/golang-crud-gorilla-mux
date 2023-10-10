package main

import (
	"fmt"
	"golang-crud/config"
	"golang-crud/routes"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()
	r := mux.NewRouter()
	routes.RoutesIndex(r)

	log.Printf("Server running on http://localhost:%v", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
}
