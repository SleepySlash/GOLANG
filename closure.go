package main

import (
	"fmt"
)

func closure() func() int{
	count:=0
	return func()int{
		count++;
		return count
	}
}

func main(){
	caller := closure()
	fmt.Println(caller())
	fmt.Println(caller())
	fmt.Println(caller())
	anotherCaller := closure()
	fmt.Println(anotherCaller())
}