package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("you have chosen a random number between 1 to 100")
	fmt.Println("can  you guess it ?")
	fmt.Println(target)
}
