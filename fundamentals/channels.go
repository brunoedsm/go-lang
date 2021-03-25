package main

import(
	"fmt"
)

// Function to attribute string len to the channel
func StrLen(s string,c chan int){
	c <- len(s)
}

func main(){

	/*Create channel*/
	c := make(chan int)

	//call function passing channel
	go StrLen("EJ 2",c)
	go StrLen("Free Hearthbreaker",c)
	go StrLen("Driving Towards the Daylight",c)
	go StrLen("Physical Grafitti",c)
	go StrLen("Part II",c)

	x, y, z, w, a := <-c, <-c, <-c, <-c, <-c
		fmt.Println(x,y,z,w,a)
	
}