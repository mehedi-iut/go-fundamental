package main

import (
	"struct_example/user"
	"fmt"

)



func main(){
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	 if err != nil{
		 fmt.Println(err)
		 return
		}
	appUser.OutputUserDetails()
	appUser.ClearUserDetails()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
