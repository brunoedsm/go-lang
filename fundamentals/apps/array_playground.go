package main

import(
	"fmt"
)

func main()  {

	var arr []int

	arr = append(arr,1)
	arr = append(arr,3)
	arr = append(arr,5)
	arr = append(arr,7)


	fmt.Println("Array created: ",arr)
	fmt.Println("Array length: ",len(arr))
	fmt.Println("Array el in 3: ",arr[3])
	fmt.Println("Array el start in 1: ",arr[1 : 0])
}