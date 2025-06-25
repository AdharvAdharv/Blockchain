package main

import "fmt"

func main() {
	
	Ages := map[string] int{"Sarath":32,"Vignesh":27,"Adharv":20}

	fmt.Println(Ages)

	for name,age := range Ages{
		if age<21 {
			fmt.Printf("%v Age is not greater than 21 \n",name)
		}else{
			fmt.Printf("%v Age is greater than 21 \n",name)
		}
	}
}