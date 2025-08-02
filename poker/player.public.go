package poker

type Player struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Stack int     `json:"stack"`
	Bet   int     `json:"bet"`
	Cards [2]Card `json:"cards"`
}

type PlayerInfo struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Stack int    `json:"stack"`
}

func NewPlayer(id string, name string, initialStack int) *Player {
	return &Player{
		Id:    id,
		Name:  name,
		Stack: initialStack,
	}
}
