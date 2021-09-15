package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"go/src/question-answer/helpers"
	"net/http"
	"github.com/gorilla/handlers"
	"os"
	"go/src/question-answer/router"
)


var db, err = helpers.GetDb()

func main() {
	//close DB after Main function's end
	if err != nil {
		panic(err)
	}

	helpers.InitEnvVal("env")

	//database.Migrate()
	//database.Seed()
	//println("Server started on port 5432")

	http.ListenAndServe(":5432", handlers.LoggingHandler(os.Stdout, router.Router()))

	defer db.Close()
}
