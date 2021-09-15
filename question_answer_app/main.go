package main

import (
	"question_answer_app/handler"
	"question_answer_app/storage/postgres"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
)

func main() {

	newDbString :=newDBFromConfig()
    store ,err := postgres.NewStorage(newDbString)

      if err != nil {
	     log.Fatal(err)
       }
  decoder := schema.NewDecoder()
  decoder.IgnoreUnknownKeys(true)
  session := sessions.NewCookieStore([]byte("my_secret"))

  r,err := handler.NewServer(store, decoder, session)

    if err != nil {
	    log.Println("handler not found")
 }


 srv := &http.Server{
  Handler:      r,
  Addr:         "127.0.0.1:8080",
  WriteTimeout: 15 * time.Second,
  ReadTimeout:  15 * time.Second,
 }
 
 log.Fatal(srv.ListenAndServe())

}

func newDBFromConfig() string {
	dbParams := " " + "user=postgres"
	dbParams += " " + "host=localhost"
	dbParams += " " + "port=5432"
	dbParams += " " + "dbname=dbnew"
	dbParams += " " + "password=password"
	dbParams += " " + "sslmode=disable"
   
	return dbParams
   }
   

