package main

import "fmt"

func main() {
	num:=20

	if num>=0 && num<=10 {
		fmt.Printf(" %v is between 0 and 10  \n",num)
	}else if num >10 && num<=20 {
		fmt.Printf("%v is between 11 and 20 ",num)
	}else{
		fmt.Printf("%v is more than 20",num )
	}
}