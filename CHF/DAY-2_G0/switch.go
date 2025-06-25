package main

import (
	"fmt"
	"time"

)

func main() {
	Day:=1

	switch Day{
	case 1:
		fmt.Println("it is Monday")
	case 2:
		fmt.Println("it is Tuesday")	
	case 3:
		fmt.Println("it is Wednesday")
	default:
		fmt.Println("it is weekend")	
	}

	now :=time.Now()

	switch {
	case now.Hour() <12:
		fmt.Println("Good Morning")
	case now.Hour() <16:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	
	}
}