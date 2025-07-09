package poker

import "errors"

type players struct {
	players      map[string]*player
	playersSlice []*player
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

func (p *players) getPlayersSliceCopy() []Player {
	playersCopy := make([]Player, 0, len(p.players))
	for _, player := range p.players {
		playersCopy = append(playersCopy, player.getPlayerCopy())
	}
	return playersCopy
}

func (p *players) getNextBettingPlayer(currentPlayer *player) *player {
	for i, player := range p.playersSlice {
		if player.id == currentPlayer.id {
			nextIndex := (i + 1) % len(p.playersSlice)
			nextPlayer := p.playersSlice[nextIndex]
			if nextPlayer.hasFolded() {
				return p.getNextBettingPlayer(nextPlayer)
			}
			return nextPlayer
		}
	}
	return nil
}
