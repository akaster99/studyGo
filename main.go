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
		fmt.Println("Add Success")
	}
	error = dictionary.Update("hello", "hello is greeting")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Update Success")
	}
	error = dictionary.Delete("hello")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Delete Success")
	}
	definition, error := dictionary.Search("hello")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(definition)
	}

}
