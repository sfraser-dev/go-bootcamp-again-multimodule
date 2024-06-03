package odd_even

import (
	"fmt"
	"math/rand"
)

func OddEven() {
	var mySlice = []int{}

	fmt.Println("\nOddEven")
	for i := 0; i < 10; i += 1 {
		mySlice = append(mySlice, rand.Intn(100))
		fmt.Printf("mySlice = %v\n", mySlice)
	}
}
