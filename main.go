package main

import "fmt"

func main() {
	fmt.Println("Hello!!")
	cards := newDeck()
	hand, remainingHand := deal(cards, 5)
	hand.print()
	fmt.Println("The remaining cards!!")
	remainingHand.print()
	fmt.Println(remainingHand.toString())
	hand.saveToFile("First")
}
