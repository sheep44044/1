package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func main() {
	num1 := 5
	num2 := 7
	sum := add(num1, num2)
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)
}
