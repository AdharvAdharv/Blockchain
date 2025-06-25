package main

import "fmt"

func main() {
	
	fmt.Println("Defer Statement LIFO")

	defer fmt.Println("This is first Statement")
	fmt.Println("This is second Statement")
	defer fmt.Println("This is Third Statement")
	fmt.Println("This is fourth Statement")
}