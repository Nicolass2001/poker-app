package poker

import (
	"math/rand"
)

// Card represents a card with value and suit
type card struct {
	value string
	suit  string
}

type cardsByPlayer struct {
	player  *player
	cardOne card
	cardTwo card
}

type deck struct {
	cards          []card
	communityCards []card
	cardsByPlayers []cardsByPlayer
}

func newDeck() *deck {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}

	cards := make([]card, 0, len(suits)*len(values))
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, card{value: value, suit: suit})
		}
	}

	communityCards := make([]card, 5)

	deck := &deck{
		cards:          cards,
		communityCards: communityCards,
	}
	deck.shuffle()
	return deck
}

func (d *deck) shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
}

func (d *deck) draw() card {
	card := d.cards[0]
	d.cards = d.cards[1:] // Remove the drawn card from the deck
	return card
}

func (d *deck) dealCardsToPlayers(players players) {
	playersSlice := players.playersSlice
	cardsByPlayers := make([]cardsByPlayer, len(playersSlice))
	for i, player := range playersSlice {
		cardsByPlayers[i] = cardsByPlayer{
			player:  player,
			cardOne: d.draw(),
			cardTwo: d.draw(),
		}
		player.cards = &cardsByPlayers[i]
	}

	d.cardsByPlayers = cardsByPlayers
}

func (c *card) getCardCopy() Card {
	return Card{
		Value: c.value,
		Suit:  c.suit,
	}
}

func (d *deck) getComunityCardsCopy() []Card {
	communityCardsCopy := make([]Card, len(d.communityCards))
	for i, card := range d.communityCards {
		communityCardsCopy[i] = card.getCardCopy()
	}
	return communityCardsCopy
}
