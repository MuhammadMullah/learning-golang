package main

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five Of Diamonds"
}

//////// LESSONS LEARNED ///////

// * In Golang, whenever you define a function, you have top the the compiler the return type
// of the value of the function as illustrated on line 11 of this file.
