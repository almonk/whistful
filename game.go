package main

import (
	"fmt"
)

type player struct {
	number int
	hand   []card
}

func newGame() {
	freshDeck := deck.generate(deck{})
	// var discardedCards []card
	var remainingCards = len(freshDeck.cards)
	const cardsDealtEachRound = 7

	p1 := player{number: 1}
	p2 := player{number: 2}

	// Deal out the cards
	for i := 0; i < cardsDealtEachRound; i++ {
		p1.hand = append(p1.hand, freshDeck.cards[i])
		println("Player 1 gets", freshDeck.cards[i].friendlyName())
	}

	remainingCards = remainingCards - cardsDealtEachRound

	// Deal to player 2

	for i := 7; i < cardsDealtEachRound*2; i++ {
		p2.hand = append(p1.hand, freshDeck.cards[i])
		println("Player 2 gets", freshDeck.cards[i].friendlyName())
	}

	remainingCards = remainingCards - cardsDealtEachRound
	fmt.Println(remainingCards, "cards remaining")

	fmt.Println("Trumps is", freshDeck.cards[cardsDealtEachRound].suit)
}
