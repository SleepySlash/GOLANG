// Implement a function that counts the number of occurrences of each word in a slice of strings and returns a map where the keys are the words and the values are the counts.

package main

import (
	"fmt"
)

func function(nums...string) map[string]int{
	m := make(map[string]int)
	for _,i :=range nums{
		count,present := m[i]
		if(present){
			m[i] = count+1
		}else{
			m[i] = 1
		}
	}
	return m
}

func main(){
	slice := make([]string,0)
	slice  = append(slice,"Rose", "is" ,"a", "rose", "is", "a", "rose", "is", "a", "rose")
	// m := make(map[string]int, 0)
	m := function(slice...)
	
	for i,j:=range m{
		fmt.Println(i,"  ",j)
	}
}