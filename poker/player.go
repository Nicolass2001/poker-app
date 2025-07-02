package poker

import "errors"

// Player represents a player
type Player struct {
	Id    string
	Name  string
	Stack int
	Bet   *BetsByPlayer
	Cards *CardsByPlayer
}

type Players struct {
	players      map[string]*Player
	playersSlice []*Player
}

// NewPlayer creates a new player with a given name and an initial stack
func NewPlayer(id string, name string, initialStack int) *Player {
	return &Player{
		Id:    id,
		Name:  name,
		Stack: initialStack,
	}
}

func newPlayers() Players {
	return Players{
		players:      make(map[string]*Player),
		playersSlice: []*Player{},
	}
}

func (p *Players) addPlayer(player *Player) error {
	if player == nil {
		return errors.New("player cannot be nil")
	}
	if _, exists := p.players[player.Id]; exists {
		return errors.New("player already exists")
	}
	p.players[player.Id] = player
	p.playersSlice = append(p.playersSlice, player)
	return nil
}

func (p *Players) getPlayersSliceCopy() []*Player {
	players := make([]*Player, 0, len(p.players))
	for _, player := range p.players {
		players = append(players, player)
	}
	return players
}
