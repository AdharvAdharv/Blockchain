package main

import "fmt"

func main() {
    fmt.Println("Arrays")
    var fruits [10] string
    fruits[3]="Mango"
    fruits[1]="Apple"
    fmt.Println(fruits[3],fruits)
	 numbers :=[10] int {1,3,4,5,5,6,6}
     fmt.Println(numbers)

     fmt.Println("Slices")
     random := numbers[2:6]
     fmt.Println(random)

     var ages [] int
     ages =append(ages, 9,6,3,4)
     fmt.Println(ages)
}