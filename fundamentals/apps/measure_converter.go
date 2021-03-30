package main

import(
	"fmt"
	"os"
	"strconv"
)

func main(){

	//program call arguments (slice)
	if len(os.Args) < 3 {
		fmt.Println("Use args: measure_converter {values} {unity}")
		os.Exit(1)
	}

	/*
		slice -> dynamic array
		:= -> create a variable and attribute in real time
		[i] -> index item of array
		[i : j] -> index item of array by position (i) 
	*/

	unityOrigin := os.Args[len(os.Args) - 1]

	valuesOrigin := os.Args[1 : len(os.Args) - 1]

	var targetUnity string

	if unityOrigin == "celsius" {
		targetUnity = "fahrenheit"
	}else if unityOrigin == "quilometros"{
		targetUnity = "milhas"
	}else{
		fmt.Println("Unknown unity")
	}

	for i,v := range valuesOrigin{
		sourceValue, err := strconv.ParseFloat(v,64)
		if err != nil{
			fmt.Printf("Incorrect value %s in %d position\n",v,i)
			os.Exit(1)
		}
	
	var targetValue float64

	if unityOrigin == "celsius"{
		targetValue = sourceValue * 1.8 + 32
	}else{
		targetValue = sourceValue / 1.60934
	}

	fmt.Printf("Operation result: %.2f in %s\n",targetValue,targetUnity)	

	
	}
	
}