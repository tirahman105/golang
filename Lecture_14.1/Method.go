package main

import "fmt"

type Doctor struct{

	Name string
	Education string
	Age int
	Salary float32
}


	//method


	func (d Doctor)getName() string {
	return d.Name
	}


	func (d Doctor)getEducation() string {
	return d.Education
	}

	func (d Doctor)getAge() int {
	return d.Age
	}

	func (d Doctor)getSalary() float32 {
	return d.Salary
	}

func main(){


	//তিন ভাবে struct এর  variable declare করা যায় 

	// পদ্ধতি ১ (এটাই সবচেয়ে ভালো পদ্ধতি বাকি ২টাও জানতে হবে) 



	var d Doctor 
	d.Name="Tareq"
	d.Education="BDS(Bachelor of Dental Surgery)"
	d.Age=29
	d.Salary= 50000.00
 

	
	// পদ্ধতি ২

	// var d=Doctor{Name:"Tareq", Education: "BDS",Age: 29, Salary:50000.00,}
	

	// পদ্ধতি ৩

	//var d=Doctor{"Tareq", "BDS", 29, 50000.00,} 

	//	fmt.Println(d)

	fmt.Println(d.getName())
	fmt.Println(d.getEducation())
}

















