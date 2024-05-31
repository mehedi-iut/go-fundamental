package main

import "fmt"

func main(){
	age := 29
	var agePointer *int
	agePointer = &age
	fmt.Println("Age: ", *agePointer)

	editAdultYears(&age)
	fmt.Println(age)
}

func editAdultYears(age *int) {
	*age = *age - 18
}
