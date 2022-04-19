package app

import (
	"fmt"
	"log"
	"net/http"
	"stocks/config"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	config.NewConfig()
	dbClient := config.NewDbConnection()
	fmt.Println(dbClient)
	log.Fatal(http.ListenAndServe(config.Cfg.Host+":"+config.Cfg.Port, router))
}
