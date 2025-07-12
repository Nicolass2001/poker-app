package poker

// handRank represents the ranking of a poker hand.
type handRank int

const (
	highCard handRank = iota
	onePair
	twoPair
	threeOfAKind
	straight
	flush
	fullHouse
	fourOfAKind
	straightFlush
	royalFlush
)

func (r handRank) String() string {
	switch r {
	case royalFlush:
		return "Royal Flush"
	case straightFlush:
		return "Straight Flush"
	case fourOfAKind:
		return "Four of a Kind"
	case fullHouse:
		return "Full House"
	case flush:
		return "Flush"
	case straight:
		return "Straight"
	case threeOfAKind:
		return "Three of a Kind"
	case twoPair:
		return "Two Pair"
	case onePair:
		return "One Pair"
	case highCard:
		return "High Card"
	default:
		return "Unknown"
	}
}
