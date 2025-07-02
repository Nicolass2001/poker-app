package poker

type BetsByPlayer struct {
	Player *Player
	Bet    int
}

type Bets struct {
	SmallBlindAmount int
	BigBlindAmount   int
	SmallBlindPlayer *Player
	BigBlindPlayer   *Player
	BetsByPlayers    []BetsByPlayer
}

func newBets(smallBlindAmount int, bigBlindAmount int) *Bets {
	bets := &Bets{
		SmallBlindAmount: smallBlindAmount,
		BigBlindAmount:   bigBlindAmount,
	}
	return bets
}

func (b *Bets) initializeBets(players Players) {
	playerSlice := players.playersSlice
	b.SmallBlindPlayer = playerSlice[0]
	b.BigBlindPlayer = playerSlice[1]
	b.BetsByPlayers = make([]BetsByPlayer, len(playerSlice))
	for i, player := range playerSlice {
		b.BetsByPlayers[i] = BetsByPlayer{
			Player: player,
			Bet:    0,
		}
		player.Bet = &b.BetsByPlayers[i]
	}
}

func (b *Bets) setBlinds() {
	b.SmallBlindPlayer.Bet.Bet = b.SmallBlindAmount
	b.BigBlindPlayer.Bet.Bet = b.BigBlindAmount
}
