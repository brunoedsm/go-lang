package main

import(
	"fmt"
)

type Card struct{
	appUserId string
	externalRef string
}

func(c Card)Save(){
	fmt.Println(c.appUserId)
	fmt.Println(c.externalRef)
}

type ICard interface{
	Save()
}

func CreateOrUpdate(class ICard){
	class.Save()
}

func main(){

	var cards = make([]ICard,0) 
	
	var c1 = new(Card)
	c1.appUserId = "INTEG1234"
	c1.externalRef = "INTEG1234-CP1"

	var c2 = new(Card)
	c2.appUserId = "QUALIF1234"
	c2.externalRef = "QUALIF1234-CP1"
	
	cards = append(cards,c1)
	cards = append(cards,c2)

	for idx, val := range cards{
		fmt.Println(idx,val)
		CreateOrUpdate(val)
	}
}
