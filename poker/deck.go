package poker

import (
	"math/rand"
)

// Card represents a card with value and suit
type Card struct {
	Value string
	Suit  string
}

type Deck struct {
	Cards []Card
}

func newDeck() *Deck {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}

	cards := make([]Card, 0, len(suits)*len(values))
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, Card{Value: value, Suit: suit})
		}
	}

	deck := &Deck{Cards: cards}
	deck.shuffle()
	return deck
}

func (d *Deck) shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

func (d *Deck) draw() Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:] // Remove the drawn card from the deck
	return card
}