package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	
)
   type Codemen struct{
	   SignatureCount int
	   Signatures []string
   }
   func check(err error) {
	   if err !=nil {
		   log.Fatal(err)
	   }
   }
   func getString(fileName string) []string {
	   var lines []string
	   file,err := os.Open(fileName)
	   if os.IsNotExist(err){
		   return nil
	   }
   
       check(err)
       defer file.Close()
       scanner :=bufio.NewScanner(file)
       for scanner.Scan() {
	   lines = append(lines,scanner.Text())
         }
       check(scanner.Err())
	   return lines
   }

	func viewHandler(writer http.ResponseWriter,request *http.Request){
		signatures := getString("signature.txt")
		html, err := template.ParseFiles("view.html")
		check (err)
		codemen := Codemen{
			SignatureCount : len(signatures),
			Signatures:      signatures,
		}
		err = html.Execute(writer,codemen)
		check (err)
		
	}
	func newHandler(writer http.ResponseWriter,request *http.Request){
		html, err :=template.ParseFiles("new.html")
		check(err)
		err=html.Execute(writer,nil)
		check(err)
	}
	func createHandler(writer http.ResponseWriter, request *http.Request){
		signature :=request.FormValue("signature")
		options :=os.O_WRONLY | os.O_APPEND | os.O_CREATE
		file,err := os.OpenFile("signature.txt",options,os.FileMode(0660))
		check(err)
		_, err = fmt.Fprintln(file,signature)
		check(err)
		err = file.Close()
		check(err)
		http.Redirect(writer,request,"/codemen",http.StatusFound)
	}
func main() {
	
		http.HandleFunc("/codemen",viewHandler)
		http.HandleFunc("/codemen/new",newHandler)
		http.HandleFunc("/codemen/create",createHandler)
		err :=http.ListenAndServe("localhost:8080",nil)
		log.Fatal(err)
}