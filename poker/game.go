package poker

import (
	"errors"
)

// Game represents the state of the game
type Game struct {
	Deck           *Deck
	Players        map[string]Player
	CommunityCards []Card
}

func NewGame(playersIds []string, playerNames []string, initialStacks []int) (*Game, error) {
	if len(playersIds) != len(playerNames) || len(playersIds) != len(initialStacks) {
		return nil, errors.New("playersIds, playerNames, and initialStacks must have the same length")
	}
	players := make(map[string]Player)
	for i := range playersIds {
		if playersIds[i] == "" || playerNames[i] == "" {
			return nil, errors.New("player ID and name cannot be empty")
		}
		if initialStacks[i] <= 0 {
			return nil, errors.New("initial stack must be greater than zero")
		}
		players[playersIds[i]] = NewPlayer(playerNames[i], initialStacks[i])
	}
	if len(players) < 2 {
		return nil, errors.New("at least two players are required to start a game")
	}

	return &Game{
		Deck:    newDeck(),
		Players: players,
	}, nil
}

// GetPlayers returns a slice of all players in the game
func (g *Game) GetPlayers() []Player {
	players := make([]Player, 0, len(g.Players))
	for _, player := range g.Players {
		players = append(players, player)
	}
	return players
}

// Deal deals two cards to each player
func (g *Game) DealCards() {
	for playerId, player := range g.Players {
		player.Cards = []Card{g.Deck.draw(), g.Deck.draw()}
		g.Players[playerId] = player
	}
}

func (g *Game) SetBets(bets map[string]int) error {
	for playerId, bet := range bets {
		err := g.setBetByPlayer(playerId, bet)
		if err != nil {
			return err
		}
	}
	// continue with the game logic after setting bets
	err := g.nextPhase()
	if err != nil {
		return err
	}
	return nil
}

// SetBetByPlayer sets the bet for a player by their ID
func (g *Game) setBetByPlayer(playerId string, bet int) error {
	player, exists := g.Players[playerId]
	if !exists {
		return errors.New("player not found: " + playerId)
	}
	err := player.SetBet(bet)
	if err != nil {
		return err
	}
	g.Players[playerId] = player
	return nil
}

func (g *Game) nextPhase() error {
	if len(g.CommunityCards) == 0 {
		g.flop()
	} else if len(g.CommunityCards) == 3 {
		g.turn()
	} else if len(g.CommunityCards) == 4 {
		g.river()
	} else if len(g.CommunityCards) == 5 {
		// All community cards have been dealt, the game can proceed to showdown or end
	} else {
		return errors.New("invalid number of community cards: " + string(rune(len(g.CommunityCards))))
	}
	return nil
}

func (g *Game) flop() {
	for range 3 {
		g.CommunityCards = append(g.CommunityCards, g.Deck.draw())
	}
}

func (g *Game) turn() {
	g.CommunityCards = append(g.CommunityCards, g.Deck.draw())
}

func (g *Game) river() {
	g.CommunityCards = append(g.CommunityCards, g.Deck.draw())
}