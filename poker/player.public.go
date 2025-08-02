package poker

type Player struct {
	Id    string
	Name  string
	Stack int
	Bet   int
	Cards [2]Card
}

type PlayerInfo struct {
	Id    string
	Name  string
	Stack int
}

func NewPlayer(id string, name string, initialStack int) *Player {
	return &Player{
		Id:    id,
		Name:  name,
		Stack: initialStack,
	}
}
