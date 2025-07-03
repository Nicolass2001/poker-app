package poker

type betsByPlayer struct {
	player *player
	bet    int
}

type bets struct {
	smallBlindAmount int
	bigBlindAmount   int
	smallBlindPlayer *player
	bigBlindPlayer   *player
	betsByPlayers    []betsByPlayer
}

func newBets(smallBlindAmount int, bigBlindAmount int) *bets {
	bets := &bets{
		smallBlindAmount: smallBlindAmount,
		bigBlindAmount:   bigBlindAmount,
	}
	return bets
}

func (b *bets) initializeBets(players players) {
	playerSlice := players.playersSlice
	b.smallBlindPlayer = playerSlice[0]
	b.bigBlindPlayer = playerSlice[1]
	b.betsByPlayers = make([]betsByPlayer, len(playerSlice))
	for i, player := range playerSlice {
		b.betsByPlayers[i] = betsByPlayer{
			player: player,
			bet:    0,
		}
		player.bet = &b.betsByPlayers[i]
	}
}

func (b *bets) setBlinds() {
	b.smallBlindPlayer.bet.bet = b.smallBlindAmount
	b.bigBlindPlayer.bet.bet = b.bigBlindAmount
}
