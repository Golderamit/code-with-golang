package main

import (
	keyboard "/home/codemen-06/go/pkg/mod/gopkg.in"
	"fmt"
	"log"
)

func main() {
	fmt.Print("Enter a grade:")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade > 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of ", grade, "is", status)
}
