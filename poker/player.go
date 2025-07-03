package poker

import "errors"

// Player represents a player
type player struct {
	id    string
	name  string
	stack int
	bet   *betsByPlayer
	cards *cardsByPlayer
}

type players struct {
	players      map[string]*player
	playersSlice []*player
}

// NewPlayer creates a new player with a given name and an initial stack
func newPlayer(p *Player) *player {
	return &player{
		id:    p.Id,
		name:  p.Name,
		stack: p.Stack,
	}
}

func newPlayers() players {
	return players{
		players:      make(map[string]*player),
		playersSlice: []*player{},
	}
}

func (p *players) addPlayer(player *Player) error {
	if player == nil {
		return errors.New("player cannot be nil")
	}
	if _, exists := p.players[player.Id]; exists {
		return errors.New("player already exists")
	}
	newPlayer := newPlayer(player)
	p.players[player.Id] = newPlayer
	p.playersSlice = append(p.playersSlice, newPlayer)
	return nil
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

func (p *players) getPlayersSliceCopy() []Player {
	playersCopy := make([]Player, 0, len(p.players))
	for _, player := range p.players {
		playersCopy = append(playersCopy, player.getPlayerCopy())
	}
	return playersCopy
}
