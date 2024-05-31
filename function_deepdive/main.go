package main

import "fmt"

type transformfn func(int)int

func main(){
	numbers := []int{1,2,3,4}
	doubled := transformNumber(&numbers, double)
	trippled := transformNumber(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(trippled)

}

func transformNumber(numbers *[]int, tansform transformfn) []int{
	dNumbers := []int{}

	for _, val := range *numbers{
		dNumbers = append(dNumbers, tansform(val))
	}

	return dNumbers
}

func double(x int) int {
	return x*2
}

func triple(x int) int {
	return x*3
}