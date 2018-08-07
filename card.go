// The model for a single playing card in Whist
//
// Ranking is
// A (14)
// K (13)
// Q (12)
// J (11)
// 10
// ..
// 1

package main

import (
	"fmt"
	"strconv"
)

type card struct {
	rank int
	suit string
}

func (card card) friendlyName() string {
	var cardName string

	switch card.rank {
	case 11:
		cardName = "Jack"
	case 12:
		cardName = "Queen"
	case 13:
		cardName = "King"
	case 14:
		cardName = "Ace"
	default:
		cardName = strconv.Itoa(card.rank)
	}

	return fmt.Sprintf("The %s of %s", cardName, card.suit)
}
