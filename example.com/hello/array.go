package main

import (
	"fmt"
	"time"
)

func main() {
	var notes [7]string
	notes[0] = "do"
	notes[1] = "se"
	notes[2] = "mi"
	notes[3] = "li"
	notes[4] = "pa"
	notes[5] = "ni"
	notes[6] = "fa"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(notes[2])
	fmt.Println(notes[3])
	var dates [3]time.Time
	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(1447920000, 0)
	dates[2] = time.Unix(1508632200, 0)
	fmt.Println(dates[0])
	fmt.Println(dates[1])
	fmt.Println(dates[2])
	// array literal
	var notess [7]string = [7]string{"do", "se", "mi", "li", "pa", "ni", "fa"}
	fmt.Println(notess[5], notess[4], notess[3])
	var primes [5]int = [5]int{4, 7, 9, 6, 25}
	fmt.Println(primes[0], primes[4], primes[3])
	var numbers [3]int
	numbers[0] = 42
	numbers[2] = 108
	var letters = [3]string{"a", "b", "c"}
	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
	fmt.Println(letters[2])
	fmt.Println(letters[0])
	fmt.Println(letters[1])
	//fmt package
	var note [7]string = [7]string{"do", "se", "mi", "li", "pa", "ni", "fa"}
	var prime [5]int = [5]int{4, 7, 9, 6, 25}
	fmt.Println(note)
	fmt.Println(prime)
	//%#v verbs used
	fmt.Printf("%#v\n", note)
	fmt.Printf("%#v\n", prime)
	//for loop

	notee := [7]string{"do", "se", "mi", "li", "pa", "ni", "fa"}
	fmt.Println(len(notee))
	for i := 0; i < len(notee); i++ {
		fmt.Println(i, notee[i])
	}
	for index, note := range notee {
		fmt.Println(index, note)
	}
	// blank identifier
	for _, note := range notee {
		fmt.Println(note)
	}
	for index, _ := range notee {
		fmt.Println(index)
	}
}
