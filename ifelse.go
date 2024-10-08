//Write a function that takes an integer as input and returns "even" if it's even, and "odd" if it's odd.

package main

import "fmt"

func evenOrOdd(num int) string {
	if num%2 == 0 {
		return "even"
	}
	return "odd"
}

func main() {
	fmt.Println("2 is ",evenOrOdd(2)) // even
	fmt.Println("3 is ",evenOrOdd(3)) // odd
}