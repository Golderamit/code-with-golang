package main

import (
	"net/http"
)




func Home(w http.ResponseWriter, r *http.Request){
	/* fmt.Fprintf(w,"this is the home page ") */
	renderTemplate(w,"home-page.html")

}
func About(w http.ResponseWriter, r *http.Request){
	/* sum:=AddValues(2,2)
	fmt.Fprintf(w,fmt.Sprintf("this is about page and addvalues for 2+2 is %d",sum)) */
	renderTemplate(w,"about-page.html")
}
/* func AddValues(x,y int) int {
	var sum int
	sum = x + y
	return sum
} */

/* func Divide(w http.ResponseWriter, r *http.Request){
f,err := dividevalues(120.0, 7.0)
  if err != nil {
	  fmt.Fprintf(w,"can not divide by 0")
  }
  fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f",120.0,7.0,f))
}

func dividevalues(x,y float32) (float32, error){
	if y<=0{
		fmt.Printf("cn not divide by 0")
		return 0,nil
	}
	result:=x/y
	return result,nil
} */

