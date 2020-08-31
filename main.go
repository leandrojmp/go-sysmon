package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leandrojmp/go-sysmon/config"
	"github.com/leandrojmp/go-sysmon/routes"
)

func handleRequests() {

	apiRouter := mux.NewRouter().StrictSlash(true)

	apiRouter.HandleFunc("/", routes.LandingPage)
	apiRouter.HandleFunc("/netstat/{port}", routes.ReturnSinglePort)
	apiRouter.HandleFunc("/netstat", routes.ReturnAllPorts)

	log.Fatal(http.ListenAndServe(config.Configuration.ListenAddress, apiRouter))
}

func main() {
	config.LoadConfig("config.json")
	handleRequests()
}
