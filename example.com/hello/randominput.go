package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("you have to chosen a random number between 1 to 100")
	fmt.Println("can you guess ?")
	fmt.Println(target)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Make a guess:   ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	if guess < target {
		fmt.Println("Oops,sorry  your guess is Low ...")
	} else if guess > target {
		fmt.Println("Oops,sorry your guess is High:  ")

	}
}
