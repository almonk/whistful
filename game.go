package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

import tl "github.com/JoelOtter/termloop"

type player struct {
	number int
	hand   []card
}

var freshDeck deck
var remainingCards int
var numberOfTurns int
var numberOfCardsDealtTotal int
var p1 player
var p2 player

type MovingText struct {
	*tl.Text
}

func newGame() {
	// var discardedCards []card
	freshDeck = deck.generate(deck{})
	remainingCards = len(freshDeck.cards)
	p1 = player{number: 1}
	p2 = player{number: 2}

	// Start the game by dealing 7 cards to each player
	// dealHeadToPlayer(p2, 7)

	// // dealHeadToPlayer(p1, 1)
	// fmt.Println(p1.hand)
	// playerPlaysATrick(p1, getPlayerInput())
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorGreen,
		Ch: 'v',
	})

	game := tl.NewGame()
	game.Screen().SetLevel(level)
	game.Screen().AddEntity(tl.NewText(0, 0, "WHIST", tl.ColorBlack, tl.ColorYellow))
	game.Screen().AddEntity(&MovingText{tl.NewText(0, 1, "AN @ALMONK GAME", tl.ColorBlack, tl.ColorGreen)})
	game.SetDebugOn(true)

	game.Start()

	dealHeadToPlayer(game, p1, 7)
}

func dealHeadToPlayer(game *tl.Game, p player, numberOfCardsToDeal int) {

	for i := 0; i < numberOfCardsToDeal; i++ {
		if numberOfCardsDealtTotal <= 51 {
			// Add cards from the deck to
			if p.number == 1 {
				p1.hand = append(p1.hand, freshDeck.cards[numberOfCardsDealtTotal])
				fmt.Println(fmt.Sprintf("Player %d gets %s", p1.number, freshDeck.cards[numberOfCardsDealtTotal].friendlyName()))
			} else {
				p2.hand = append(p1.hand, freshDeck.cards[numberOfCardsDealtTotal])
				fmt.Println(fmt.Sprintf("Player %d gets %s", p2.number, freshDeck.cards[numberOfCardsDealtTotal].friendlyName()))
			}

			drawNextCardX := 26 * numberOfCardsDealtTotal
			game.Screen().AddEntity(tl.NewRectangle(2, 3, drawNextCardX, 12, tl.RgbTo256Color(255, 255, 255)))

			// Increment the number of cards dealt
			numberOfCardsDealtTotal++
		} else {
			fmt.Println("There are no more cards that can be dealt")
		}
	}
}

func playerPlaysATrick(p player, cardInHand int) {
	// A player lays a card down
	fmt.Println(fmt.Sprintf("Player %d is playing a trick...", p.number))
	fmt.Println(fmt.Sprintf("Player %d played %s...", p.number, p.hand[cardInHand].friendlyName()))
	time.Sleep(1 * time.Second)
	playerPlaysATrick(getNextPlayer(p), 1)
}

func getNextPlayer(p player) player {
	var nextPlayer player
	if p.number == 1 {
		nextPlayer = p2
	} else {
		nextPlayer = p1
	}
	return nextPlayer
}

func pickTrumps() {
	fmt.Println("Trumps is", freshDeck.cards[numberOfCardsDealtTotal+1].suit)
}
