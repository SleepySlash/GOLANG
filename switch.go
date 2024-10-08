// Create a function that takes a weekday number (1 for Monday, 2 for Tuesday, etc.) and returns the corresponding weekday name using a switch statement.

package main

import (
	"fmt"
)
func weekdayName(weekday int) string{
	switch weekday{
	case 1: return "Monday"
	case 2: return "Tuesday"
	case 3: return "Wednesday"
	case 4: return "Thursday"
	case 5: return "Friday"
	case 6: return "Saturday"
	case 7: return "Sunday"
	default: return "Invalid weekday"
	}
}

func main(){
	for i:=1;i<=7;i++{
		fmt.Println("day ",i," is ",weekdayName(i))
	}
}