package main

import(
	"fmt"
)

func main(){
	
	/*Primitives*/
	var x = "Hello World"
	fmt.Println(x)

	z := int(42)
	fmt.Println(z)

	/*Slices and maps*/

	var s = make([]string,0)
	var m = make(map[string]string)

	s = append(s,"string one")
	s = append(s,"string two")
	m["1972"] = "Bad Co"
	m["1973"] = "Straight Shooter"

	fmt.Println(s)
	fmt.Println(s[1])
	
	fmt.Println(m)
	var title = m["1972"]
	fmt.Println(title)

}