package main

import "fmt"

func main() {
	numbers := [3]float64{54.5, 45.0, 36.25}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: % 0.2f\n ", sum/sampleCount)

	numberss := [6]int{3, 16, -2, 10, 23, 12}
	for i, numberss := range numberss {
		if numberss >= 10 && numberss < 20 {
			fmt.Println(i, numberss)
		}

	}

}
