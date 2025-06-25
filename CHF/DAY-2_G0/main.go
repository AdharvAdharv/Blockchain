package main

import "fmt"

func main() {
	 fmt.Println("Welcome to Go programming Language")
	 var name string
	 fmt.Print("Enter Your Name :")
	 fmt.Scan(&name)
	 fmt.Printf("Hi %s Welcome to the company \n",name)
}