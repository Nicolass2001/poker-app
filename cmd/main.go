package main

import (
	"fmt"
	"poker-app/poker"
	"strings"
)

func showPlayers(players []poker.Player) {
	fmt.Printf("%-15s %-30s %-10s\n", "Player", "Cards", "Bet")
	fmt.Println(strings.Repeat("-", 60))
	for _, player := range players {
		card1 := player.Cards[0].String()
		card2 := player.Cards[1].String()
		fmt.Printf("%-15s %-30s %-10d\n", player.Name, fmt.Sprintf("%s, %s", card1, card2), player.Bet)
	}
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println()
}

func showCommunityCards(cards []poker.Card) {
	fmt.Printf("Community Cards: ")
	for _, card := range cards {
		fmt.Printf("%s of %s, ", card.Value, card.Suit)
	}
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
	game.MakeAction(poker.ActionCall, 0)

	println(" - After Player 2 calls the big blind:")
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

	game.MakeAction(poker.ActionRaise, 500)

	println(" - After Player 1 raises:")
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

	game.MakeAction(poker.ActionCall, 0)

	println(" - After Player 2 calls the raise:")
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

	game.MakeAction(poker.ActionCheck, 0)

	println(" - Should be in the Flop state now:")
	println(" - After Player 2 checks:")
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

	game.MakeAction(poker.ActionCheck, 0)

	println(" - After Player 1 checks:")
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

	game.MakeAction(poker.ActionCheck, 0)

	println(" - Should be in the Turn state now:")
	println(" - After Player 2 checks again:")
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

	game.MakeAction(poker.ActionAllIn, 0)

	println(" - After Player 1 goes all-in:")
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

	game.MakeAction(poker.ActionCall, 0)

	println(" - After Player 2 calls the all-in:")
	showPlayers(game.GetPlayers())
	showCommunityCards(game.GetCommunityCards())

}
