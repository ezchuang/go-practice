package main

import "fmt"

// 1
// func main() {
// 	fmt.Println("Hi there!")
// }

// 2
// func main() {
// 	var card string = "Five of Diamonds"

// 	fmt.Println(card)
// }

// 3
// func main() {
// 	card := "Ace of Spades"
// 	card = "Five of Diamonds"

// 	fmt.Println(card)
// }

// 4
func main() {
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
