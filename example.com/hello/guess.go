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

	reader := bufio.NewReader(os.Stdin)

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("you have", 10-guesses, "guesses left...")
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
		} else {
			success = true
			fmt.Println("Good jobs! You guessed it!")
			break
		}
	}
	if !success {
		fmt.Println("sorry!! You did not guess my number!it is :", target)
	}

}
