package poker

type hand struct {
	cards       []card
	highestCard card
	handRank    handRank
}

func newHandWithCards(cards []card) *hand {
	highestCard := evaluateHighestCard(cards)
	handRank := evaluateHand(cards)
	return &hand{
		cards:       cards,
		highestCard: highestCard,
		handRank:    handRank,
	}
}

func evaluateHighestCard(cards []card) card {
	highest := cards[0]
	for _, c := range cards {
		if c.value > highest.value {
			highest = c
		}
	}
	return highest
}

func evaluateHand([]card) handRank {
	// TODO: Implement hand evaluation logic
	return highCard
}

func (h *hand) compareHands(otherHand *hand) int {
	if h.getHandRank() > otherHand.getHandRank() {
		return 1
	} else if h.getHandRank() < otherHand.getHandRank() {
		return -1
	}
	// If ranks are equal, compare the highest cards
	if h.highestCard.compareCards(otherHand.highestCard) == 1 {
		return 1
	}
	if h.highestCard.compareCards(otherHand.highestCard) == -1 {
		return -1
	}
	return 0
}

func (h *hand) getHandRank() handRank {
	return h.handRank
}

func (h *hand) string() string {
	return h.handRank.String() + " with highest card " + h.highestCard.value + " of " + h.highestCard.suit
}
