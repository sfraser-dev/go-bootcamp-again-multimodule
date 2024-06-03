package main

import (
	"bitbucket.org/weebucket/gobootcampagain_multimodule/structs/reusable"
	"fmt"
)

func main () {
	var p1 my_structs.Person
	p1.FirstName = "Joe"
	p1.LastName = "Bloggs"
	fmt.Printf("p1 = %v\n", p1)
}