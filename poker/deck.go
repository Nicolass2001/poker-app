package poker

import (
	"math/rand"
)

type cardsByPlayer struct {
	player   *player
	cardOne  card
	cardTwo  card
	bestHand *hand
}

type deck struct {
	cards          []card
	communityCards []card
	cardsByPlayers []cardsByPlayer
}

func newDeck() *deck {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	cards := make([]card, 0, len(suits)*len(values))
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, card{value: value, suit: suit})
		}
	}

	communityCards := make([]card, 0, 5)

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
	d.cards = d.cards[1:]
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

func (d *deck) getComunityCardsCopy() []Card {
	communityCardsCopy := make([]Card, len(d.communityCards))
	for i, card := range d.communityCards {
		communityCardsCopy[i] = card.getCardCopy()
	}
	return communityCardsCopy
}

func (d *deck) flop() {
	d.communityCards = append(d.communityCards, d.draw(), d.draw(), d.draw())
}

func (d *deck) turn() {
	d.communityCards = append(d.communityCards, d.draw())
}

func (d *deck) river() {
	d.communityCards = append(d.communityCards, d.draw())
}

func (d *deck) calculateWinners() []*player {
	winners := make([]*player, 0)
	for i := range d.cardsByPlayers {
		posibleCards := append(d.communityCards, d.cardsByPlayers[i].cardOne, d.cardsByPlayers[i].cardTwo)
		d.cardsByPlayers[i].bestHand = evaluateBestHand(posibleCards)
		winners = updateWinners(winners, d.cardsByPlayers[i].player)
	}
	return winners
}

func updateWinners(winners []*player, playerToCompare *player) []*player {
	if len(winners) == 0 {
		return append(winners, playerToCompare)
	}

	bestHandByNow := winners[0].cards.bestHand
	playerHand := playerToCompare.cards.bestHand
	handsComparison := playerHand.compareHands(bestHandByNow)

	if handsComparison == 1 {
		return []*player{playerToCompare}
	}
	if handsComparison == 0 {
		winners = append(winners, playerToCompare)
	}
	return winners
}

func evaluateBestHand(cards []card) *hand {
	combinations := combinationsOf5(cards)
	bestHand := newHandWithCards(combinations[0])
	for _, combination := range combinations {
		hand := newHandWithCards(combination)
		if hand.compareHands(bestHand) == 1 {
			bestHand = hand
		}
	}
	return bestHand
}

func combinationsOf5(input []card) [][]card {
	var result [][]card
	var comb []card = make([]card, 5)

	var combine func(start, depth int)
	combine = func(start, depth int) {
		if depth == 5 {
			// make a copy of comb and add to result
			tmp := make([]card, 5)
			copy(tmp, comb)
			result = append(result, tmp)
			return
		}
		for i := start; i < len(input); i++ {
			comb[depth] = input[i]
			combine(i+1, depth+1)
		}
	}

	combine(0, 0)
	return result
}
