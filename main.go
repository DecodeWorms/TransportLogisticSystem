package main

import (
	"Tsystem/config"
	"Tsystem/handlers"
	"Tsystem/storage"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var client *storage.Client
var store storage.DataStore
var seaStore storage.Store

func init() {
	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	var err error
	var db *gorm.DB

	cfg := config.Config{
		DatabaseHost:     host,
		DatabaseUsername: user,
		DatabaseName:     dbname,
		DatabasePort:     port,
	}

	initctx := context.Background()
	client = storage.NewConnection(initctx, cfg, db, err)
	store = storage.NewDataStore(initctx, client)
	seaStore = storage.NewStore(initctx, client)

}

func main() {

	var router = mux.NewRouter()

	var road = handlers.NewRoadHandlers(&store)
	var sea = handlers.NewSea(&seaStore)

	router.HandleFunc("/road/create", road.Create).Methods("POST")
	router.HandleFunc("/road/roads", road.Roads).Methods("GET")
	router.HandleFunc("/road/road/{tracker}", road.Road).Methods("GET")
	router.HandleFunc("/road/changeLocation", road.UpdateLocation).Methods("PUT")
	router.HandleFunc("/road//delete/{tracker}", road.Delete).Methods("DELETE")

	router.HandleFunc("/sea/create", sea.Create).Methods("POST")
	router.HandleFunc("/sea/seas", sea.Seas).Methods("GET")
	router.HandleFunc("/sea/sea/{tracker}", sea.Sea).Methods("GET")
	router.HandleFunc("/sea/update", sea.Update).Methods("PUT")
	router.HandleFunc("/sea/delete", sea.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println(client)

}
