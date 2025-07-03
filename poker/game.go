package poker

import (
	"errors"
)

// Game represents the state of the game
type game struct {
	deck      *deck
	players   players
	bets      *bets
	gameState gameState
}

// NewGame creates a new game instance with the specified small and big blind amounts
func NewGame(smallBlindAmount int, bigBlindAmount int) (*game, error) {
	bets := newBets(smallBlindAmount, bigBlindAmount)
	deck := newDeck()
	players := newPlayers()

	return &game{
		deck:      deck,
		bets:      bets,
		players:   players,
		gameState: stateWaitingForPlayers,
	}, nil
}

// AddPlayer adds a player to the game
func (g *game) AddPlayer(player *Player) error {
	if g.gameState != stateWaitingForPlayers {
		return errors.New("game has already started, cannot add players")
	}
	return g.players.addPlayer(player)
}

// StartGame initializes the game state, deals cards to players and sets the blinds
func (g *game) StartGame() error {
	// Check if the game is in the waiting state
	if g.gameState != stateWaitingForPlayers {
		return errors.New("game is already started or in progress")
	}
	// Check if there are enough players to start the game
	if len(g.players.players) < 2 {
		return errors.New("not enough players to start the game")
	}
	g.bets.initializeBets(g.players)
	g.bets.setBlinds()
	g.deck.dealCardsToPlayers(g.players)
	g.gameState = statePreFlop
	return nil
}

// GetPlayers returns a slice of all players in the game
func (g *game) GetPlayers() []Player {
	return g.players.getPlayersSliceCopy()
}

// GetComunityCards returns the community cards
func (g *game) GetCommunityCards() []Card {
	return g.deck.getComunityCardsCopy()
}
