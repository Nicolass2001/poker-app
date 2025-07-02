package poker

type GameState int

const (
	StateWaitingForPlayers GameState = iota
	StatePreFlop
	StateFlop
	StateTurn
	StateRiver
	StateShowdown
	StateHandOver
)

func (s GameState) String() string {
	switch s {
	case StateWaitingForPlayers:
		return "Waiting for Players"
	case StatePreFlop:
		return "Pre-Flop"
	case StateFlop:
		return "Flop"
	case StateTurn:
		return "Turn"
	case StateRiver:
		return "River"
	case StateShowdown:
		return "Showdown"
	case StateHandOver:
		return "Hand Over"
	default:
		return "Unknown State"
	}
}
