package poker

import (
	"errors"
)

// Game represents the state of the game
type Game struct {
	Deck      *Deck
	Players   Players
	Bets      *Bets
	GameState GameState
}

func NewGame(smallBlindAmount int, bigBlindAmount int) (*Game, error) {
	bets := newBets(smallBlindAmount, bigBlindAmount)
	deck := newDeck()
	players := newPlayers()

	return &Game{
		Deck:      deck,
		Bets:      bets,
		Players:   players,
		GameState: StateWaitingForPlayers,
	}, nil
}

// AddPlayer adds a player to the game
func (g *Game) AddPlayer(player *Player) error {
	if g.GameState != StateWaitingForPlayers {
		return errors.New("game has already started, cannot add players")
	}
	return g.Players.addPlayer(player)
}

// StartGame initializes the game state, deals cards to players and sets the blinds
func (g *Game) StartGame() error {
	// Check if the game is in the waiting state
	if g.GameState != StateWaitingForPlayers {
		return errors.New("game is already started or in progress")
	}
	// Check if there are enough players to start the game
	if len(g.Players.players) < 2 {
		return errors.New("not enough players to start the game")
	}
	g.Bets.initializeBets(g.Players)
	g.Bets.setBlinds()
	g.Deck.dealCardsToPlayers(g.Players)
	g.GameState = StatePreFlop
	return nil
}

// GetPlayers returns a slice of all players in the game
func (g *Game) GetPlayers() []*Player {
	return g.Players.getPlayersSliceCopy()
}

// GetComunityCards returns the community cards
func (g *Game) GetCommunityCards() []Card {
	return g.Deck.getComunityCardsCopy()
}
