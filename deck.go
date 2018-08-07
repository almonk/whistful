// Modelling a deck of 52 cards
package main

import (
	"math/rand"
	"time"
)

type deck struct {
	cards []card
}

func (deck) generate() deck {
	suits := [4]string{"Hearts", "Spades", "Clubs", "Diamonds"}
	currentSuit := 0

	var newDeck = deck{}

	// Generate all the cards in the deck
	for i := 0; i < 4; i++ {
		for j := 2; j < 15; j++ {
			newDeck.cards = append(newDeck.cards, card{rank: j, suit: suits[currentSuit]})
		}
		currentSuit++
	}

	// Now shuffle them...
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(newDeck.cards), func(i, j int) { newDeck.cards[i], newDeck.cards[j] = newDeck.cards[j], newDeck.cards[i] })

	return newDeck
}

func removeCardAt(s deck, index int) []card {
	return append(s.cards[:index], s.cards[index+1:]...)
}
