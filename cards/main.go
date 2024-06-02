package main

import (
	"fmt"
)

func main() {
	// var card string = "ace of spades" 	// formal declaration
	//card := "ace of spades" 				// type inference declaration
	/*
		card := getOneCard()
		fmt.Printf("card = %s; card=%q\n", card, card) // string, double quoted str. snippet "fmpf"
		cardsWee := deck{"Two of Clubs", getOneCard()}    // deck is type slice of string
		cards := append(cardsWee, "Five of Hearts")    // argument cards_wee not changed
	*/
	/*
		// for loop
		for i := 0; i < len(cards); i += 1 {
		 	fmt.Println(cards[i]) // snippet "fp" or "fmpl"
		 }
	*/
	/*
		// for loop range keyword
		for i, card := range cards { // range is a keyword in Go, returns arr/slice ind and copy of value
			fmt.Printf("i = %3d, card = %15s", i, card)
			fmt.Println("")
		}
	*/

	var cards deck = newDeck()            // new deck of cards
	fmt.Println("\nprint the whole deck") // receiver functions are LIKE "OOP" in Go
	cards.printWholeDeck()
	cards.shuffleDeck() // shuffle the deck
	fmt.Println("\nprint the shuffled deck")
	cards.printWholeDeck()

	fmt.Println("\ndeal two hands of 5 cards")
	hand1, remainingCards := deal(cards, 5)
	fmt.Println("hand1:")
	hand1.printWholeDeck()
	fmt.Println("hand2:")
	hand2, remainingCards := deal(remainingCards, 5)
	hand2.printWholeDeck()

	fmt.Println("\nbyte slice interlude")
	hiStr := "Hi there"
	bs := []byte(hiStr)
	fmt.Println(hiStr) // hi there
	fmt.Println(bs)    // [72 105 32 116 104 101 114 101]

	fmt.Println("\nhand 1 to single string")
	h1 := toSingleString(hand1)
	fmt.Println(h1)

	fmt.Println("\nhand 1 reconstructed to string slice")
	hand11 := fromSingleString(h1, ":")
	hand11.printWholeDeck()

	fmt.Println("\nwrite the whole deck to file")
	cards.writeToFile()

	fmt.Println("\nread the whole deck from file")
	cards.readFromFile()
	cards.printWholeDeck()
}
