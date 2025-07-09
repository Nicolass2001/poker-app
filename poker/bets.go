package poker

import "errors"

type betByPlayer struct {
	player   *player
	bet      int
	isFolded bool
}

type bets struct {
	smallBlindAmount    int
	bigBlindAmount      int
	smallBlindPlayer    *player
	bigBlindPlayer      *player
	bettingPlayer       *player
	highestBet          int
	bigBlindPlayerBeted bool
	betsByPlayers       []betByPlayer
}

func newBets(smallBlindAmount int, bigBlindAmount int) *bets {
	bets := &bets{
		smallBlindAmount: smallBlindAmount,
		bigBlindAmount:   bigBlindAmount,
		highestBet:       bigBlindAmount,
	}
	return bets
}

func (b *bets) initializeBets(players players) {
	playerSlice := players.playersSlice
	b.bigBlindPlayer = playerSlice[0]
	b.smallBlindPlayer = playerSlice[1]
	b.betsByPlayers = make([]betByPlayer, len(playerSlice))
	for i, player := range playerSlice {
		b.betsByPlayers[i] = betByPlayer{
			player:   player,
			bet:      0,
			isFolded: false,
		}
		player.bet = &b.betsByPlayers[i]
	}
}

func (b *bets) setBlinds() {
	b.smallBlindPlayer.raiseBet(b.smallBlindAmount)
	b.bigBlindPlayer.raiseBet(b.bigBlindAmount)
	b.bettingPlayer = b.smallBlindPlayer
	b.bigBlindPlayerBeted = false
}

func (b *bets) getBettingPlayer() *player {
	return b.bettingPlayer
}

func (b *bets) setBettingPlayer(player *player) {
	b.bettingPlayer = player
}

func (b *bets) playerAction(player *player, action Action, amount int) error {
	currentBet := b.highestBet

	switch action {
	case ActionCheck:
		if player.bet.bet < currentBet {
			return errors.New("cannot check when not matched")
		}

	case ActionCall:
		diff := currentBet - player.bet.bet
		if diff <= 0 {
			return errors.New("nothing to call")
		}
		player.raiseBet(diff)

	case ActionRaise:
		if amount <= 0 {
			return errors.New("invalid amount for action")
		}
		player.raiseBet(amount)

	case ActionAllIn:
		player.raiseBet(player.stack)

	case ActionFold:
		player.bet.isFolded = true

	default:
		return errors.New("invalid action")
	}

	if !player.hasFolded() && player.bet.bet > b.highestBet {
		b.highestBet = player.bet.bet
	}

	if player.id == b.bigBlindPlayer.id {
		b.bigBlindPlayerBeted = true
	}
	return nil
}

func (b *bets) keepBetting() bool {
	if !b.bigBlindPlayerBeted {
		return true
	}
	for _, bet := range b.betsByPlayers {
		if !bet.isFolded && bet.bet < b.highestBet {
			return true
		}
	}
	return false
}

func (b *bets) newBettingRound() {
	b.bigBlindPlayerBeted = false
	b.bettingPlayer = b.smallBlindPlayer
}
