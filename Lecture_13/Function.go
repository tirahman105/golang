package main

import ("fmt")

//example_1
/*
func add(x int, y int) int {

	r := x+y

	return r 
*/

//example_2


/*
func add(x, y int) int {

	r := x+y

	return r 

}

*/

/*
func add(x, y int) (r int) {

	r = x+y

	return r 

}
*/


func add(x, y int) (r int) {

	r = x+y

	return 

}



func rectangle(l int, b int) (area int, parameter int){
		parameter = 2*(l+b)
		area = l*b
		return //Return statement without specific variable name  
}


func main(){

	a,p :=rectangle(10,10)
	fmt.Println(a,p)

}



