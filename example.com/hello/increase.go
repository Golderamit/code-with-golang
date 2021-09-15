package main
import"fmt"

func main(){
    var b int=15
    var a int
    numbers:=[6]int{1,2,3,5} 

	for a:=0;a<10;a++{
		fmt.Printf("value of a: %b\n",a)

	}
	for a<b{
		a++
		fmt.Printf("value of a: %b\n",a)
	}
	for i,x:=range numbers{
		fmt.Printf("value of x= %d\n at %d\n",x,i)
	}
}