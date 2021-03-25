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

type IFriend interface{
	SayHello()
}

func Greet(f IFriend){
	f.SayHello()
}

func main(){

	var qa = new(Person)
	qa.Name = "John Doe"
	qa.Age = 36
	Greet(qa)

}


