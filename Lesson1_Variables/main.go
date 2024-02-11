package main

import "fmt"

func main() {
	//Declearing Variables
	var a int = 3
	// var b string = 's'
	var b string = "svjj"
	// pp := 12

	fmt.Println("Hello",a);
	y := fmt.Sprintf("Hello %d",1+1)
	fmt.Println("Hello",b);
	fmt.Printf("Hello %s",b);
	fmt.Printf("Hello %s\n",y);

   if 1==1 {
	fmt.Println("good")
   }
	//In Golang, string interpolation refers to the process of embedding variables or expressions within a string
	// fmt.Printf("hello %v Abebe",pp)
	// fmt.Printf("hello %v",b)


}