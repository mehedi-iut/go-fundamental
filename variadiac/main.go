package main

import "fmt"

func main(){
	numbers := []int{10,12,34,56}

	total := sumNumber(10,12,34,56)
	anotherTotal := sumNumber(numbers...)
	
	fmt.Println("sum: ", total)
	fmt.Println(anotherTotal)

}

func sumNumber(numbers ...int) int {
	result := 0

	for _, val := range numbers{
		result += val
	}

	return result
}