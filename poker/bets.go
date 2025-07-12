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

func (b *bets) setNextBettingPlayer() {
	b.bettingPlayer = b.getNextBettingPlayer(b.bettingPlayer)
}

func (b *bets) getNextBettingPlayer(player *player) *player {
	for i, bet := range b.betsByPlayers {
		if bet.player.id == player.id {
			nextIndex := (i + 1) % len(b.betsByPlayers)
			nextBet := b.betsByPlayers[nextIndex]
			if nextBet.isFolded {
				return b.getNextBettingPlayer(nextBet.player)
			}
			return nextBet.player
		}
	}
	return nil
}

func (b *bets) playerAction(action Action, amount int) error {
	amountToRaise, err := b.calculateAmountToRaise(action, amount)
	if err != nil {
		return err
	}

	bettingPlayer := b.bettingPlayer
	bettingPlayer.raiseBet(amountToRaise)

	b.highestBet = max(b.highestBet, bettingPlayer.bet.bet)

	if bettingPlayer.id == b.bigBlindPlayer.id {
		b.bigBlindPlayerBeted = true
	}
	return nil
}

func (b *bets) calculateAmountToRaise(action Action, amount int) (int, error) {
	currentBet := b.highestBet
	bettingPlayer := b.bettingPlayer
	amountToRaise := amount

	switch action {
	case ActionCheck:
		if bettingPlayer.bet.bet < currentBet {
			return 0, errors.New("cannot check when not matched")
		}
		amountToRaise = 0

	case ActionCall:
		diff := currentBet - bettingPlayer.bet.bet
		if diff <= 0 {
			return 0, errors.New("nothing to call")
		}
		amountToRaise = diff

	case ActionRaise:
		if amount <= 0 {
			return 0, errors.New("invalid amount for action")
		}
		amountToRaise = amount

	case ActionAllIn:
		amountToRaise = bettingPlayer.stack

	case ActionFold:
		bettingPlayer.bet.isFolded = true
		amountToRaise = 0

	default:
		return 0, errors.New("invalid action")
	}

	return amountToRaise, nil
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

func (b *bets) distributeWinnings(winners []*player) {
	totalPot := 0
	for _, bet := range b.betsByPlayers {
		totalPot += bet.bet
	}

	winningsPerWinner := totalPot / len(winners)
	for _, winner := range winners {
		winner.stack += winningsPerWinner
	}
}
