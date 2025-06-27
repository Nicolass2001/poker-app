package poker

import (
	"errors"
)

// Player represents a player
type Player struct {
	Id string
	Name  string
	Cards []Card
	Stack int
	Bet   int
}

// NewPlayer creates a new player with a given name and an initial stack
func NewPlayer(name string, initialStack int) Player {
	return Player{
		Name:  name,
		Cards: []Card{},
		Stack: initialStack,
		Bet:   0,
	}
}

// SetBet sets the bet for the player and deducts it from their stack
func (p *Player) SetBet(bet int) error {
	if bet > p.Stack {
		return errors.New("bet exceeds player's stack: " + p.Name + " has " + string(rune(p.Stack)) + " but tried to bet " + string(rune(bet)))
	}
	p.Bet += bet
	p.Stack -= bet
	return nil
}

