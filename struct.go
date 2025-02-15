package main

import "fmt"
  
type person struct{
	name string
	age  int
	lastname string
}
     func main () {
		var nikhil person
		nikhil.name= "nikhil"
		nikhil.age = 23
		nikhil.lastname= "pachpute"
		fmt.Println("nikhil person;", nikhil)

}
