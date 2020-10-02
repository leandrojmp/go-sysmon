package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		config.CreateLoggers()
		config.LoadConfig("config.json")
		config.InfoLogger.Print("application started")
		handleRequests()
	}()
	<-exitSignal
	config.InfoLogger.Print("application stopped")
}
