package poker

type gameState int

const (
	stateWaitingForPlayers gameState = iota
	statePreFlop
	stateFlop
	stateTurn
	stateRiver
	stateShowdown
	stateHandOver
)

func (s gameState) String() string {
	switch s {
	case stateWaitingForPlayers:
		return "Waiting for Players"
	case statePreFlop:
		return "Pre-Flop"
	case stateFlop:
		return "Flop"
	case stateTurn:
		return "Turn"
	case stateRiver:
		return "River"
	case stateShowdown:
		return "Showdown"
	case stateHandOver:
		return "Hand Over"
	default:
		return "Unknown State"
	}
}

func (s gameState) bettingState() bool {
	switch s {
	case statePreFlop, stateFlop, stateTurn, stateRiver:
		return true
	default:
		return false
	}
}

func (s gameState) nextState() gameState {
	switch s {
	case stateWaitingForPlayers:
		return statePreFlop
	case statePreFlop:
		return stateFlop
	case stateFlop:
		return stateTurn
	case stateTurn:
		return stateRiver
	case stateRiver:
		return stateShowdown
	case stateShowdown:
		return stateHandOver
	default:
		return stateWaitingForPlayers
	}
}
