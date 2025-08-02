package poker

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
}

func (c *Card) String() string {
	return c.Value + " of " + c.Suit
}
