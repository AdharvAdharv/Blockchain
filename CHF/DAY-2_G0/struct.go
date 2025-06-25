package main

import "fmt"

type User struct{
	Name string
	Age int
	Email string
		
}
func main() {
	
	user1 := User{"Lavanya",25 ,"lavanya@gmail.com"}
	user2 :=User{"Vineesh", 27,"vineesh@gmail.com"}

	fmt.Println(user1,user2)
	fmt.Println(user1.Email)
}