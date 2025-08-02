package poker

import (
	"errors"
)

// Game represents the state of the game
type Game struct {
	deck      *deck
	players   players
	bets      *bets
	gameState gameState
}

// NewGame creates a new game instance with the specified small and big blind amounts
func NewGame(smallBlindAmount int, bigBlindAmount int) *Game {
	bets := newBets(smallBlindAmount, bigBlindAmount)
	deck := newDeck()
	players := newPlayers()

	return &Game{
		deck:      deck,
		bets:      bets,
		players:   players,
		gameState: stateWaitingForPlayers,
	}
}

// AddPlayer adds a player to the game
func (g *Game) AddPlayer(player *Player) error {
	if g.gameState != stateWaitingForPlayers {
		return errors.New("game has already started, cannot add players")
	}
	return g.players.addPlayer(player)
}

// StartGame initializes the game state, deals cards to players and sets the blinds
func (g *Game) StartGame() error {
	if g.gameState != stateWaitingForPlayers {
		return errors.New("game is already started or in progress")
	}
	if len(g.players.players) < 2 {
		return errors.New("not enough players to start the game")
	}
	g.bets.initializeBets(g.players)
	g.bets.setBlinds()
	g.deck.dealCardsToPlayers(g.players)
	g.gameState = g.gameState.nextState()
	return nil
}

// GetPlayerById returns the player with the specified ID
func (g *Game) GetPlayerById(id string) (Player, error) {
	return g.players.getPlayerById(id)
}

// GetPlayers returns a slice of all players in the game
func (g *Game) GetPlayers() []Player {
	return g.players.getPlayersSliceCopy()
}

// GetPlayersInfo returns a slice of PlayerInfo for all players in the game (PlayerInfo: id, name, chips)
func (g *Game) GetPlayersInfo() []PlayerInfo {
	return g.players.getPlayersInfoSliceCopy()
}

// GetComunityCards returns the community cards
func (g *Game) GetCommunityCards() []Card {
	return g.deck.getComunityCardsCopy()
}

// GetCurrentPlayer returns the PlayerInfo of the player whose turn it is to act
func (g *Game) GetCurrentPlayer() PlayerInfo {
	return g.bets.getBettingPlayer().getPlayerInfoCopy()
}

// GetGameState returns the current game state
func (g *Game) GetGameState() string {
	return g.gameState.String()
}

// MakeAction allows the current player to perform an action (check, call, raise, fold, allin) with a specified amount
func (g *Game) MakeAction(action Action, amount int) error {
	if !g.gameState.bettingState() {
		return errors.New("game is not in a valid state for actions")
	}

	err := g.bets.playerAction(action, amount)
	if err != nil {
		return err
	}

	if g.bets.keepBetting() {
		g.bets.setNextBettingPlayer()
		return nil
	}

	g.nextBettingGameState()
	return nil
}

func (g *Game) nextBettingGameState() {
	g.bets.newBettingRound()

	switch g.gameState {
	case statePreFlop:
		g.deck.flop()
	case stateFlop:
		g.deck.turn()
	case stateTurn:
		g.deck.river()
	}

	g.gameState = g.gameState.nextState()

	if g.gameState == stateShowdown {
		g.showdownLogic()
	}
}

func (g *Game) showdownLogic() {
	winners := g.deck.calculateWinners()
	println("Winners of the showdown:")
	for _, winner := range winners {
		println(" -", winner.name, "with hand:", winner.cards.bestHand.string())
	}
	g.bets.distributeWinnings(winners)
}
