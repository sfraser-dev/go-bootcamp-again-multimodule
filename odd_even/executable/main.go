package main

import (
	"github.com/sfraser-dev/go-bootcamp-again-multimodule/odd_even/reusable"
)

func main() {
	mySlice1, mySlice2 := odd_even.OddEven()
	odd_even.PrintOddAndEven(mySlice1, mySlice2)
}