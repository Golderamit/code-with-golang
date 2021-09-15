package main

import (
	"fmt"
)

func main() {
	var width, height, area float64
	width = 4.2
	height = 3.0
	area = width * height
	fmt.Printf("%.2f liter needed\n", area/10.0)
	width = 5.2
	height = 3.5
	area = width * height
	fmt.Printf("%.2f liter needed\n", area/10.0)

	//Formatting value widths

	fmt.Printf("%12s    | %s\n", "product", "liter needed")
	fmt.Println("---------------------------------------")
	fmt.Printf("%12s    |%2d\n", "wall width", 50)
	fmt.Printf("%12s    |%2d\n", "wall height", 5)
	fmt.Printf("%12s    |%2d\n", "Total area", 99)
	//quick demonstration of various width values in action:
	fmt.Printf("%%7.3f: %7.3f\n", 12.3456)
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456)
	fmt.Printf("%%.1f: %.1f\n", 12.3456)
	fmt.Printf("%%.2f: %.2f\n", 12.3456)
	fmt.Printf("%%.3f: %.3f\n", 12.3456)
}
