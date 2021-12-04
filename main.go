package main

import "fmt"

func superAdd(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func canIDrink(age int) bool {
	if age < 18 {
		return false
	}
	return true
}
func main() {
	super := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(super)
}
