package main
import "fmt"

func inRange(min float64,max float64,numbers ...float64) []float64{
	var result []float64
	for_,number:=range numbers{
		if number >= min && number <= max {
			result=append(result, number)
		}
	}
	return result
	
}
func main(){
	fmt.Println(inRange(1,100,-12.5,3.2,0,50,103.50))
	fmt.Println(inRange(10,-10,4.1,12,-12,-5.2)
}