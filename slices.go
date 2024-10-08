package main

import (
	"fmt"
	"slices"
)

func max(num ...int) int {
	var max int
	for _, i := range num {
		if max < i {
			max = i
		}
	}
	return max
}

func main() {
	var slice = make([]int, 5)
	slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,1343,424,42345,4485)
	fmt.Println("using funciton ",max(slice...))
	fmt.Println("using method ",slices.Max(slice))

}