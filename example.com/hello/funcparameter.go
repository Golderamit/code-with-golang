package main

import "fmt"

func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liter needed ..\n", area/10.0)
}
func main() {
	paintNeeded(2.3, 3.0)
	paintNeeded(5.0, 3.5)
	paintNeeded(3.5, 3.7)
}
