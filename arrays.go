// Write a function that accepts an array of 5 integers and returns the sum of its elements.

package main

import (
	"fmt"
)

func sumOfArray(num ...int) int{
	var sum int
	for _,i:=range num{
		sum+=i
	}
	return sum
}

func main(){
	fmt.Println(sumOfArray(1,2,3,4,5,6,7,8,9,10))
	var arr = []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(sumOfArray(arr...))
}