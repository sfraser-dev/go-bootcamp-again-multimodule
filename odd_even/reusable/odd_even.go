package odd_even

import (
	"fmt"
	"math/rand"
)

func OddEven() {
	var mySlice = []int{}
	var myResults = []string{}

	fmt.Println("\nOddEven")
	for i := 0; i < 10; i += 1 {
		mySlice = append(mySlice, rand.Intn(100))
		fmt.Printf("mySlice = %v\n", mySlice)
	}
	for i := 0; i < 10; i += 1 {
		if mySlice[i]%2 == 0 {
			myResults = append(myResults, "E")
		} else {
			myResults = append(myResults, "O")
		}
	}
	fmt.Printf("myResults = %v\n", myResults)

}
