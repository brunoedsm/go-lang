package main

import(
	"fmt"
	"encoding/json"
)

type CustomError string
func (e CustomError) Error() string{
	return string(e)
}

type Foo struct {
	Bar string
	Baz string
}

type FooXML struct {
    Bar     string    `xml:"id,attr"`
    Baz     string    `xml:"parent>child"`
}


func main(){

	f := Foo{"Joe Junior", "Hello Shabado"}
    b, _ := json.Marshal(f)
    fmt.Println(string(b))
       json.Unmarshal(b, &f)

	g := FooXML{"Joe Manfra", "TFS"}
    fmt.Println(g)

}