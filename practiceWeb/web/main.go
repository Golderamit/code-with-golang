package main

import (
	"log"
	"net/http"
	"time"
	"practiceWeb/web/handler"
	"practiceWeb/web/storage/postgres"
	
)
func main(){

	dbString :=newDBFromConfig()
	store,err :=postgres.NewStorage(dbString)
	if err != nil{
		log.Fatal(err)
	}
	r,err :=handler.NewServer(store)
	if err != nil{
		log.Fatal("handler not found")
	}
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
func newDBFromConfig() string{
	dbParams := " " + "user=postgres"
	dbParams += " " + "host=localhost"
	dbParams += " " + "port=5435"
	dbParams += " " + "dbname=practice"
	dbParams += " " + "password=secret"
	dbParams += " " + "sslmode=disable"

	return dbParams
}