package main

import(
	"fmt"
	"net/http"
	"github.com/stacktitan/ldapauth" /*third party*/
)

func main()  {
	fmt.Println("GoGet test")
}
/*
	go get {package url} => download dependencies for code
	tree src github.com/stacktitan/ => list source codes in repo * only in Linux
	go lint {src} => repair syntax (standalone tool and can be found in go get -u golang.org/x/lint/golint)
	go vet {src} => check the code to find issues
	go test {src} => run tests in code
	go cover {src} => check test coverage
*/