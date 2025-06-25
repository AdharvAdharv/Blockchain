package main

import "fmt"


type User struct{
	FitstName string
	LastName string
	Age int
}
func main() {

	user1:=User{"Bill","Gates",66}
	user2:=User{"Richard","Branson",72}
	user3:=User{"Tom","Holland",28}

	fmt.Println(user1,user2,user3)
	
}