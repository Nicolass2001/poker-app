package poker

import "errors"

type player struct {
	id    string
	name  string
	stack int
	bet   *betByPlayer
	cards *cardsByPlayer
}

func newPlayer(p *Player) *player {
	return &player{
		id:    p.Id,
		name:  p.Name,
		stack: p.Stack,
	}
}

func (p *player) getPlayerCopy() Player {
	return Player{
		Id:    p.id,
		Name:  p.name,
		Stack: p.stack,
		Bet:   p.bet.bet,
		Cards: [2]Card{
			p.cards.cardOne.getCardCopy(),
			p.cards.cardTwo.getCardCopy(),
		},
	}
}

func (p *player) raiseBet(amount int) error {
	if amount <= 0 {
		return errors.New("raise amount must be greater than zero")
	}
	if p.stack < amount {
		return errors.New("not enough stack to raise")
	}
	p.bet.bet += amount
	p.stack -= amount
	return nil
}

func (p *player) hasFolded() bool {
	return p.bet.isFolded
}
