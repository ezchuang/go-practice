package main

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
// func main() {
// 	card := newCard()

// 	fmt.Println(card)
// }

// func newCard() string {
// 	return "Five of Diamonds"
// }

// 5
// func main() {
// 	cards := []string{"Ace of Diamonds", newCard()}
// 	cards = append(cards, "Six of Spades")

// 	for i, card := range cards {
// 		fmt.Println(i, card)
// 	}
// }

// func newCard() string {
// 	return "Five of Diamonds"
// }

// 6
func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
