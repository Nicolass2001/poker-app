package poker

type Card struct {
	Value string
	Suit  string
}

func (c *Card) String() string {
	return c.Value + " of " + c.Suit
}
