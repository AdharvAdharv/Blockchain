package main

import "fmt"

func main() {
	array := [5]int{4,7,12,9,15}
	fmt.Println("Array :",array)
	
	var sum int
	largest :=0

	for _,v := range array{
       sum+=v  

	   if v>largest {
		largest=v
		
	   }

	}
	fmt.Println("Sum of Array : ",sum) 
	fmt.Println("Largest num in array",largest)

	

}