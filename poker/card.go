package poker

import "strconv"

type card struct {
	value string
	suit  string
}

func (c *card) getCardCopy() Card {
	return Card{
		Value: c.value,
		Suit:  c.suit,
	}
}

func (c *card) compareCards(other card) int {
	thisValue, thisErr := strconv.Atoi(c.value)
	otherValue, otherErr := strconv.Atoi(other.value)
	if thisErr != nil && otherErr != nil {
		// values are letters
		return compareLetters(c.value, other.value)
	}
	if thisErr != nil {
		// this is a letter, other is a number
		return 1
	}
	if otherErr != nil {
		// other is a letter, this is a number
		return -1
	}
	if thisValue > otherValue {
		return 1
	} else if thisValue < otherValue {
		return -1
	}
	return 0
}

func compareLetters(a, b string) int {
	if a == b {
		return 0
	}
	if a == "A" ||
		(a == "K" && b != "A") ||
		(a == "Q" && b != "A" && b != "K") ||
		(a == "J" && b != "A" && b != "K" && b != "Q") {
		return 1
	}
	return -1
}
