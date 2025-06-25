package main

import "fmt"

func main() {
	
	var p *int
	i:=34
	p=i
	fmt.Println(i)
	fmt.Println(*p)
	fmt.Println(&p)
}