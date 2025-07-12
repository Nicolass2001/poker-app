package poker

type hand struct {
	cards    []card
	handRank handRank
}

func newHandWithCards(cards []card) *hand {
	cards = orderCardsByValue(cards)
	handRank := evaluateHand(cards)
	return &hand{
		cards:    cards,
		handRank: handRank,
	}
}

func orderCardsByValue(cards []card) []card {
	for i := 0; i < len(cards)-1; i++ {
		for j := 0; j < len(cards)-i-1; j++ {
			if cards[j].compareCards(cards[j+1]) == 1 {
				cards[j], cards[j+1] = cards[j+1], cards[j]
			}
		}
	}
	return cards
}

func evaluateHand(cards []card) handRank {
	if checkRoyalFlush(cards) {
		return royalFlush
	}
	if checkStraightFlush(cards) {
		return straightFlush
	}
	if checkFourOfAKind(cards) {
		return fourOfAKind
	}
	if checkFullHouse(cards) {
		return fullHouse
	}
	if checkFlush(cards) {
		return flush
	}
	if checkStraight(cards) {
		return straight
	}
	if checkThreeOfAKind(cards) {
		return threeOfAKind
	}
	if checkTwoPair(cards) {
		return twoPair
	}
	if checkOnePair(cards) {
		return onePair
	}
	return highCard
}

func checkRoyalFlush(cards []card) bool {
	if checkFlush(cards) && checkStraight(cards) {
		if cards[0].value == "A" {
			return true
		}
	}
	return false
}

func checkStraightFlush(cards []card) bool {
	if checkFlush(cards) && checkStraight(cards) {
		return true
	}
	return false
}

func checkFourOfAKind(cards []card) bool {
	for i := 0; i < len(cards)-3; i++ {
		if cards[i].value == cards[i+1].value &&
			cards[i].value == cards[i+2].value &&
			cards[i].value == cards[i+3].value {
			return true
		}
	}
	return false
}

func checkFullHouse(cards []card) bool {
	if cards[0].value == cards[1].value && cards[1].value == cards[2].value && cards[3].value == cards[4].value {
		return true
	}
	if cards[0].value == cards[1].value && cards[2].value == cards[3].value && cards[3].value == cards[4].value {
		return true
	}
	return false
}

func checkFlush(cards []card) bool {
	for _, c := range cards {
		if c.suit != cards[0].suit {
			return false
		}
	}
	return true
}

func checkStraight(cards []card) bool {
	for i := 0; i < len(cards)-1; i++ {
		if cards[i].nextValue() != cards[i+1].value {
			if i == 3 && cards[4].value == "A" && cards[0].value == "2" {
				return true
			}
			return false
		}
	}
	return true
}

func checkThreeOfAKind(cards []card) bool {
	for i := 0; i < len(cards)-2; i++ {
		if cards[i].value == cards[i+1].value && cards[i].value == cards[i+2].value {
			return true
		}
	}
	return false
}

func checkTwoPair(cards []card) bool {
	pairCount := 0
	for i := 0; i < len(cards)-1; i++ {
		if cards[i].value == cards[i+1].value {
			pairCount++
			i++ // Skip the next card since it's part of the pair
		}
	}
	return pairCount == 2
}

func checkOnePair(cards []card) bool {
	for i := 0; i < len(cards)-1; i++ {
		if cards[i].value == cards[i+1].value {
			return true
		}
	}
	return false
}

func (h *hand) compareHands(otherHand *hand) int {
	if h.getHandRank() > otherHand.getHandRank() {
		return 1
	} else if h.getHandRank() < otherHand.getHandRank() {
		return -1
	}
	// If ranks are equal, compare the highest cards
	for i := range h.cards {
		if h.cards[i].compareCards(otherHand.cards[i]) == 1 {
			return 1
		} else if h.cards[i].compareCards(otherHand.cards[i]) == -1 {
			return -1
		}
	}
	return 0
}

func (h *hand) getHandRank() handRank {
	return h.handRank
}

var suitSymbols = map[string]string{
	"Spades":   "♠",
	"Hearts":   "♥",
	"Diamonds": "♦",
	"Clubs":    "♣",
}

func (h *hand) string() string {
	return h.handRank.String() + " with cards: " + toString(h.cards)
}

func toString(cards []card) string {
	result := ""
	for _, c := range cards {
		result += c.value + " of " + suitSymbols[c.suit] + ", "
	}
	if len(result) > 0 {
		result = result[:len(result)-2] // Remove trailing comma and space
	}
	return result
}
