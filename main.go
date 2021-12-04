package main

import (
	"fmt"

	"github.com/akaster99/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First Word"}
	error := dictionary.Add("hello", "greeting")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Success")
	}
	definition, error := dictionary.Search("hello")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(definition)
	}

}
