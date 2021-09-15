package main
import"fmt"
func main(){
	
	//var mySlice []string
	var notes[]string
	notes=make([]string, 10)
	notes[0]="do"
	notes[1]="re"
	notes[2]="mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	
	//declare a slice variable
	primes:=make([]int, 10)
	primes[0]=4
	primes[1]=8
	fmt.Println(primes[1])

	letters:=[]string{"a","b","c"}
	for i:=0;i<len(letters);i++{
		fmt.Println(letters[i])
	}
	for _,letter:=range letters{
         fmt.Println(letter)
	}
	notess:=[]string{"A","M","I","T","G","O","L","D","E","R"}
	fmt.Println(notess[0],notess[1],notess[2],notess[3],notess[4],notess[5],notess[6],notess[7],notess[8],notess[9])
	primess:=[]int{28,12,1991}
	fmt.Println(primess[0],primess[1],primess[2])
	
	//slice operator
	underlyingArray:=[5]string{"A","M","I","T","G"}
	slice1:=underlyingArray[0:3]
	fmt.Println(slice1)
	underlyngArray:=[5]string{"A","M","I","T","G"}
	i,j:=1,4
	slice2:=underlyngArray[i:j]
	fmt.Println(slice2)

	slice3:=underlyngArray[2:5]
	fmt.Println(slice3)

	slice4:=underlyngArray[:3]
    fmt.Println(slice4)

	slice5:=underlyngArray[1:]
	fmt.Println(slice5)

	slice6:=underlyngArray[0:3]
	slice7:=underlyngArray[2:5]
	fmt.Println(slice6,slice7)
// Change the underlying array and the slice

	array1:=[5]string{"A","M","I","T","G"}
	slice8:=array1[0:3]
	array1[1]="X"
	fmt.Println(array1)
	fmt.Println(slice8)

	array2:=[5]string{"A","M","I","T","G"}
	slice9:=array2[2:5]
	slice9[2]="x"
	fmt.Println(array2)
	fmt.Println(slice9)

	array3:=[5]string{"A","M","I","T","G"}
	slice10:=array3[0:3]
	slice11:=array3[2:5]
	array3[2]="Co"
	fmt.Println(array3)
	fmt.Println(slice10,slice11)

//append in slice
slice:=[]string{"a","b"}
fmt.Println(slice,len(slice))
slice=append(slice, "c")
fmt.Println(slice,len(slice))
slice=append(slice, "d","e")
fmt.Println(slice,len(slice))

s1:=[]string{"s1","s1"}
s2:=append(s1,"s2","s2")
s3:=append(s2,"s3","s3")
s4:=append(s3,"s4","s4")
fmt.Println(s1,s2,s3,s4)
s4[0]="xx"
fmt.Println(s1,s2,s3,s4)
}