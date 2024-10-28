package main

import "fmt"

func returns(a int,b int) (int,int){
	return a+b,a*b
}
func main(){
	sum,product:=returns(10,5)
	fmt.Println("the sum and products of 10 and 5 are",sum,product)
	sum,product =returns(41,779)
	fmt.Println("the sum and products of 41 and 779 are",sum,product)
}