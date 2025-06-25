package main

import "fmt"

func main() {
	phonebook:=map[string]string{"Jason":"00321"}
	phonebook["Varun"]="9893"
	phonebook["Hiran"]="1234"
	fmt.Println(phonebook)
	fmt.Println(phonebook["Varun"])

}