package main

import(
	"fmt"
	"os"
	"log"
)
func main()
{
	fileInfo,err:=os.Stat("my.text")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}