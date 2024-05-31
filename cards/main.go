package main

import "fmt"

func main() {
	// var card string = "ace of spades" 	// formal declaration
	//card := "ace of spades" 				// type inference declaration
	card := newCard()
	fmt.Printf("card = %s; card=%q\n", card, card)   // string, double quoted str. snippet "fmpf"
	cards_wee := []string{"Two of Clubs", newCard()} // slice of strings
	cards := append(cards_wee, "Five of Hearts")     // argument cards_wee not changed

	// for i := 0; i < len(cards); i += 1 {
	// 	fmt.Println(cards[i]) // snippet "fp" or "fmpl"
	// }

	// range keyword
	for i, card := range cards { // range is a keyword in Go, returns arr/slice ind and copy of value
		fmt.Printf("i = %3d, card = %15s", i, card)
		fmt.Println("")
	}
}

func newCard() string {
	return "Ace of Spades"
}
