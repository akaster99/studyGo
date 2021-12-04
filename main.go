package main

import (
	"fmt"

	"github.com/akaster99/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First Word"}
	definition, error := dictionary.Search("first")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(definition)
	}
}
