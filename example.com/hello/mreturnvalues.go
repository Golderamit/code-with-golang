package main

import (
	"fmt"
	"math"
)

func floatParts(number float64) (integerPart int, fractionalPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

func main() {
	cans, remainder := floatParts(5.32598)
	fmt.Println(cans, remainder)

}
