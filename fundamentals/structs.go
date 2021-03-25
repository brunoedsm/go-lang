package main

import(
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func (p *Person) SayHello(){
	fmt.Println("Hello: ",p.Name)
	fmt.Println("Age: ",p.Age)
}

func main(){

	var qa = new(Person)
	qa.Name = "Hadhami Hadassa"
	qa.Age = 35
	qa.SayHello()
	
}


