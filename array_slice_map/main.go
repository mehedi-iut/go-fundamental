package main

import "fmt"

type Product struct {
	title string
	id int
	price float64
}

func main() {
	hobbies := [3]string{"sleeping", "eating", "learning"}
	
	fmt.Println(hobbies)
	
	firstHobby := hobbies[0]
	otherHobbies := hobbies[1:3]
	fmt.Println(firstHobby)
	fmt.Println(otherHobbies)
	
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)
	
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)
	
	
	goals := []string{"expert", "deep_dive"}
	
	fmt.Println(goals)
	
	goals[1] = "devops_expert"
	goals = append(goals, "create_operator")
	fmt.Println(goals)
	
	
	products := []Product{
		Product{title: "Apple", id: 1, price: 100.23},
		Product{title: "Samsung", id: 2, price: 23.34},
	}
	
	fmt.Println(products)
	
	products = append(products, Product{title: "One_plus", id: 3, price: 10.12})
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.


// ------------------------------------------------------------------------------------
// -----------------------------------------------------------------------------------
//func main(){
//	prices := [4]float64{10.1, 34.2, 43.2, 98.9}
//	fmt.Println(prices)
//	
//	highLightedPrice := prices[1:]
//	fmt.Println(highLightedPrice)
//	featuredPrice := highLightedPrice[:1]
//	fmt.Println(featuredPrice)
//	fmt.Println(len(featuredPrice), cap(featuredPrice))
//	
//	featuredPrice2 := featuredPrice[:3]
//	fmt.Println(featuredPrice2)
//}