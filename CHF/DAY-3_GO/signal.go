package main

import "fmt"

func main() {
	
	var light string

	fmt.Println("Enter the color of traffic light :")
	fmt.Scan(&light)

	switch light{
	case "red" :
		fmt.Println("Stop")

	case "green" :
		fmt.Println("Go")
		
	case "orange" :
		fmt.Println("Be ready to stop")	
	
	default :
		fmt.Println("Invalid Input")
	}

}