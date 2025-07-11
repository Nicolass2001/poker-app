package main

import (
	"fmt"
	"poker-app/poker"
	"strings"
)

var suitSymbols = map[string]string{
	"Spades":   "♠",
	"Hearts":   "♥",
	"Diamonds": "♦",
	"Clubs":    "♣",
}

func cardsToString(card poker.Card) string {
	suitSymbol, ok := suitSymbols[card.Suit]
	if !ok {
		suitSymbol = "?"
	}
	return fmt.Sprintf("%s%s", card.Value, suitSymbol)
}

func showPlayers(players []poker.Player) {
	fmt.Printf("%-15s %-30s %-10s %-10s\n", "Player", "Cards", "Stack", "Bet")
	fmt.Println(strings.Repeat("-", 60))
	for _, player := range players {
		card1 := cardsToString(player.Cards[0])
		card2 := cardsToString(player.Cards[1])
		fmt.Printf("%-15s %-30s %-10d %-10d\n", player.Name, fmt.Sprintf("%s, %s", card1, card2), player.Stack, player.Bet)
	}
	fmt.Println(strings.Repeat("-", 60))
}

func showCommunityCards(cards []poker.Card) {
	fmt.Printf("Community Cards: ")
	for _, card := range cards {
		fmt.Printf("%s, ", cardsToString(card))
	}
	fmt.Println()
	fmt.Println()
}

func main() {
	// Initialize a new game
	playerIds := []string{"player1", "player2"}
	playerNames := []string{"Player 1", "Player 2"}
	initialStacks := []int{10000, 20000}
	smallBlindAmount := 100
	bigBlindAmount := 200
	game, err := poker.NewGame(smallBlindAmount, bigBlindAmount)
	if err != nil {
		fmt.Println("Error creating game:", err)
		return
	}
	for i, playerId := range playerIds {
		player := poker.NewPlayer(playerId, playerNames[i], initialStacks[i])
		err := game.AddPlayer(player)
		if err != nil {
			fmt.Println("Error adding player:", err)
			return
		}
	}
	err = game.StartGame()
	if err != nil {
		fmt.Println("Error starting game:", err)
		return
	}

	// Show initial game state
	fmt.Println(" - Game started with the following players:")
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

	// Player actions

	println(" - After Player 2 calls the big blind:")
	game.MakeAction(poker.ActionCall, 0)
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

	println(" - After Player 1 raises:")
	game.MakeAction(poker.ActionRaise, 500)
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

	println(" - After Player 2 calls the raise:")
	game.MakeAction(poker.ActionCall, 0)
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

	println(" - Should be in the Flop state now:")
	println(" - After Player 2 checks:")
	game.MakeAction(poker.ActionCheck, 0)
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

	println(" - After Player 1 checks:")
	game.MakeAction(poker.ActionCheck, 0)
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

	println(" - Should be in the Turn state now:")
	println(" - After Player 2 checks again:")
	game.MakeAction(poker.ActionCheck, 0)
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

	println(" - After Player 1 goes all-in:")
	game.MakeAction(poker.ActionAllIn, 0)
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

	println(" - After Player 2 calls the all-in:")
	game.MakeAction(poker.ActionCall, 0)
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

	println(" - Should be in the River state now:")
	println(" - After Player 2 checks:")
	game.MakeAction(poker.ActionCheck, 0)
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

	println(" - After Player 1 checks:")
	game.MakeAction(poker.ActionCheck, 0)
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

}
