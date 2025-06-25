package main

import "fmt"

func main() {
	
	fmt.Println("Usual for loop")
	sum:=1
	for i:=0;i<10;i++{
		sum+=i
		fmt.Println(sum)
	}

	fmt.Println("The while loop")
	for sum<1000{
		sum+=sum
		fmt.Println(sum)
	}
	
	fmt.Println("Range Function")
	numbers:=[5]int{3,4,5,6,9}
	for k,v :=range numbers{
		fmt.Print(k)
		fmt.Println(v)
	}

}