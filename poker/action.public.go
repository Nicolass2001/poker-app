package poker

type Action int

const (
	ActionCheck Action = iota
	ActionCall
	ActionRaise
	ActionAllIn
	ActionFold
)

func (a Action) String() string {
	switch a {
	case ActionCheck:
		return "check"
	case ActionCall:
		return "call"
	case ActionRaise:
		return "raise"
	case ActionAllIn:
		return "all-in"
	case ActionFold:
		return "fold"
	default:
		return "unknown"
	}
}
