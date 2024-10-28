package main

import (
	"fmt"
)

func main(){
	fmt.Println("product of 2 and 3 is ",product(2,3))
	fmt.Println("product of 41 and 779 is ",product(41,779))
}

func product(a int,b int) int{
	return a*b
}