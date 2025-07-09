package poker

type Player struct {
	Id    string
	Name  string
	Stack int
	Bet   int
	Cards [2]Card
}

func NewPlayer(id string, name string, initialStack int) *Player {
	return &Player{
		Id:    id,
		Name:  name,
		Stack: initialStack,
	}
}
