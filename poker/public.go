package poker

type Card struct {
	Value string
	Suit  string
}

type Player struct {
	Id    string
	Name  string
	Stack int
	Bet   int
	Cards [2]Card
}

func (c *Card) String() string {
	return c.Value + " of " + c.Suit
}

func NewPlayer(id string, name string, initialStack int) *Player {
	return &Player{
		Id:    id,
		Name:  name,
		Stack: initialStack,
	}
}
