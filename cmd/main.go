package main

import (
	"server/packages/api"
	"server/packages/db"
	"server/packages/config"
	"log"
	"net/http"
)

func main() {
	// Initialise environment variables
	config.InitConfig()

	// Connects to database
	database, err := db.Connect()
	if database == nil {
		panic(err)
	} else {
		database.AutoMigrate(&db.User{}, &db.Task{})
	}

	// Initialise router
	r := api.Router(&api.Handler{DB: database})
	http.Handle("/", r)

	// Starts server
	log.Fatal(http.ListenAndServe(config.Config[config.SERVER_PORT], r))
}