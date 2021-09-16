package main

import (
	"net/http"

	"github.com/Golderamit/code-with-golang/pkg/handlers"
)
const portNumber =":8080"

func main() {
	http.HandleFunc("/",handlers.Home)
	http.HandleFunc("/about",handlers.About)
	/* http.HandleFunc("/divide",Divide) */


	_= http.ListenAndServe(portNumber,nil)
}