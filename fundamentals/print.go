package main

import(
	"fmt"
)

func main()  {
	fmt.Println("Hello World in GO")
}

/*
	go print.go => execute code in terminal
	go build print.go => create a .exe artifact to run code
	go build -ldflags "-w -s" print.go => create a .exe artifact without debug info files
	GOOS="linux" GOARCH="amd64" go build print.go => create a .exe artifact full compatible with architecture informed **linux
	file print.exe will describe the target architecture 
	go doc fmt.Println => create a output in console describing the function
*/