package main

import "fmt"


func update(a *int){
	fmt.Println("Where ",a)
	*a=*a+10
}

func main(){

var x int

fmt.Println("x value is", x)
fmt.Println("x memory address is", &x)


var y *int

fmt.Println("y value is", y)
fmt.Println("y memory address is", &y)

x= 10  // assignment
y = &x //referencing

fmt.Println("x value is", x) //accessing 
fmt.Println("y dereferancing value is", *y)


	update(&x)
	fmt.Println(x)


/*
var a string
a="Tareq"

fmt.Println(a, &a)

*/

/*
var y *int 
  
y*= 10

fmt.Println(*y)
*/

//var name *string

//fmt.Println(name)


} 

