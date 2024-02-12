package main

import "fmt"
type Awe interface{
	walk() string
	distance() int 
}
type Person struct{
	age int
	name string
}
func (p Person) walk()string{
	return "walk"
}
func (p Person) distance()int{
	return 123
}
//nested struct
type Address struct{
	address string
	city string
}
func (p Address) walk()string{
	return "Address"
}
func (p Address) distance()int{
	return 1234
}
// Interface 1
type Logger interface {
	Log(message string)
}

// Interface 2
type Messenger interface {
	Send(message string)
}

// Concrete type implementing both interfaces
type MessageService struct{}

// Implementing Logger interface
func (ms MessageService) Log(message string) {
	fmt.Println("Logging:", message)
}

// Implementing Messenger interface
func (ms MessageService) Send(message string) {
	fmt.Println("Sending:", message)
}

type Police struct{
	location string
	Address 
}
func Lo(a Awe){
	fmt.Println(a.walk())
	fmt.Println(a.distance())

}

func main(){
//Anonumous Struct
var po Police = Police{location:"Addis Ababa",Address:Address{address:"Bole",city:"gule"}}
	mystruct := struct{
		a int
		b string
	}{
		a :12,
		b:"dddd",
	}
	mymessage := MessageService{}
    person1 := Person{}
	person1.age = 12
  fmt.Println("dd",person1);
  fmt.Println("dd",mystruct);
  fmt.Println("dd",po.address);
  mymessage.Log("kl")
  Lo(person1)
}