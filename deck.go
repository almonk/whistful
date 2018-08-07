// Modelling a deck of 52 cards
package main

func generateDeck() []card {
	suits := [4]string{"Hearts", "Spades", "Clubs", "Diamonds"}
	currentSuit := 0

	var deck []card

	for i := 0; i < 4; i++ {
		for j := 2; j < 15; j++ {
			deck = append(deck, card{rank: j, suit: suits[currentSuit]})
		}
		currentSuit++
	}

	return deck
}
