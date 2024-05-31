package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

func New(firstName, lastName, birthdate string) (*User, error){
	if firstName == "" || lastName == "" || birthdate == ""{
		return nil, errors.New("firstname, lastname and birthdate is required")
	}
	return &User{
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func(u User) OutputUserDetails(){
	// u is short for (*u).firstName
	// it is only allow for struct in go
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func(u *User) ClearUserDetails(){
	u.firstName = ""
	u.lastName = ""
}