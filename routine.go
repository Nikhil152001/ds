package main

import (
	"fmt"
	"time"
)
       

func sayHello() {
	fmt.Println("hi nikhil")
    time.Sleep(5000 * time.Millisecond)
}
func sayHi() {
	fmt.Println("nikhil")
}

func main() {
	fmt.Println("learning go")
	  go sayHello()
	 go sayHi()
    time.Sleep(6000* time.Microsecond)
}
