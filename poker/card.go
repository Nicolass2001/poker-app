package poker

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
