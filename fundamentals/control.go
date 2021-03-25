package main

import(
	"fmt"
)

func main(){

	var x = 1
	var y = "none"

	if x == 1 {
		fmt.Println("X is equal to 1")
	}else{
		fmt.Println("X is not equal to 1")
	}

	switch  y {
	case "foo":
		fmt.Println("Found foo")
	case "bar":
		fmt.Println("Found bar")
	default:
		fmt.Println("Default case")
	}

	foo(x)
	foo(y)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	nums := []int{2,4,6,8}
   for idx, val := range nums {
       fmt.Println(idx, val)
   }

}

func foo(i interface{}){
	switch i.(type) {
	case int:
		fmt.Println("Found integer")
	case string:
		fmt.Println("Found string")
	default:
		fmt.Println("Unknown type")
	}
}