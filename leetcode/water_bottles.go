package main

import "fmt"


func main() {

	test := numWaterBottles(15, 5)
	fmt.Println(test)
	
}

func numWaterBottles(numBottles int, numExchange int) int {
	
	result := numBottles
	emptyBottles := numBottles

	for emptyBottles >= numExchange {

		// exchange
		numBottles = emptyBottles / numExchange
		emptyBottles = emptyBottles % numBottles

		// drink
		emptyBottles += numBottles 
		result += numBottles	
	}
	return result
}