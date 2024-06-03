package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// new type called deck of type sting slice
type deck []string

const (
	mySeperator string = ":"
)

// receiver: "extend" the functionality of the deck type via receiver
func (d deck) printWholeDeck() {
	for _, card := range d {
		fmt.Printf("%25s\n", card)
	}
}

// receiver
func (d deck) shuffleDeck() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

// receiver
func (d deck) writeToFile(filename string) {
	err := os.WriteFile(filename, []byte(toSingleString(d, mySeperator)), 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// helper functions
func readFromFile(filename string) (myStrSlice []string) {
	var myByteSlice = []byte{}
	myByteSlice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return fromSingleString(string(myByteSlice), mySeperator)
}

func newDeck() deck {
	cards := deck{}
	// creating two LISTS, suits and numbers
	suits := [4]string{"Clubs", "Diamonds", "Hearts", "Spades"}
	numbers := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range suits {
		for _, number := range numbers {
			tempStr := number + " of " + suit
			cards = append(cards, tempStr)
		}
	}
	return cards
}

func deal(d deck, handSize int) (hand deck, remainingCards deck) {
	return d[:handSize], d[handSize:]
}

func toSingleString(d deck, sep string) (theStr string) {
	// var s string = ""
	// for _, card := range d {
	// 	s = s + card + sep
	// }
	theStr = strings.Join(d, sep)
	return theStr
}

func fromSingleString(s string, sep string) (d deck) {
	d = strings.Split(s, sep)
	return d
}
