package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	app "github.com/kucinghitam13/alien/internal"
	"github.com/kucinghitam13/alien/internal/config"
)

var (
	router *httprouter.Router
)

func init() {
	log.Println("Initiating alien")

	appConfig := config.InitConfig()

	initApp(appConfig.Config)
}

func initApp(config *config.Config) {
	log.Println("Initiating service")
	service, err := app.Init(config)
	if err != nil {
		log.Panicln("[ERR] Error initiating service")
	}
	log.Println("Service has been initialized")

	log.Println("Setting endpoint")
	router = httprouter.New()
	err = service.SetEndpoint(router)
	if err != nil {
		log.Panicln("[ERR] Error setting endpoint")
	}
	log.Println("Endpoint has been set")
}

func main() {
	log.Println("Server started on: http://localhost:1981")
	log.Fatal(http.ListenAndServe(":1981", router))
}
