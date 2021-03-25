package main

import(
	"fmt"
)

// Interface
type IMusician interface{
	Who()
}

// Call
func Call(m IMusician){
	m.Who()
}

// Classes
type Bassist struct{}
	func (b *Bassist) Who(){
		fmt.Println("I'm bassist")
	}
 

type Guitarrist struct{}
	func (b *Guitarrist) Who(){
		fmt.Println("I'm guitar player")
	}

type Drummer struct{}
	func (b *Drummer) Who(){
		fmt.Println("I'm drummer")
	}


func main(){
	
	var b = new(Bassist)
	var g = new(Guitarrist)
	var d = new(Drummer)
	Call(b)
	Call(g)
	Call(d)

}