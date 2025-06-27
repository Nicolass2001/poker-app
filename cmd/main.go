package main

import (
	"fmt"
	"poker-app/poker"
	"strings"
)

// ShowPlayers displays the players and their cards
func showPlayers(players []poker.Player) {
	fmt.Printf("%-15s %-30s %-10s\n", "Player", "Cards", "Bet")
	fmt.Println(strings.Repeat("-", 60))
	for _, player := range players {
		card1 := fmt.Sprintf("%s of %s", player.Cards[0].Value, player.Cards[0].Suit)
		card2 := fmt.Sprintf("%s of %s", player.Cards[1].Value, player.Cards[1].Suit)
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
	playerNames := []string{"Player 1", "Player 2"}
	initialStacks := []int{10000, 10000}
	game, err := poker.NewGame(playerNames, playerNames, initialStacks)
	if err != nil {
		fmt.Println("Error creating game:", err)
		return
	}

	// Deal cards and show players
	game.DealCards()
	fmt.Println("Initial Hands:")
	showPlayers(game.GetPlayers())

	// Set bets pre-flop
	bets := map[string]int{
		"Player 1": 100,
		"Player 2": 150,
	}
	err = game.SetBets(bets)
	if err != nil {
		fmt.Println("Error setting bets:", err)
		return
	}
	// Show comunity cards after betting pre-flop
	showCommunityCards(game.CommunityCards)
	showPlayers(game.GetPlayers())

	// Set bets Flop
	bets = map[string]int{
		"Player 1": 200,
		"Player 2": 250,
	}
	err = game.SetBets(bets)
	if err != nil {
		fmt.Println("Error setting bets:", err)
		return
	}
	// Show community cards after betting Flop
	showCommunityCards(game.CommunityCards)
	showPlayers(game.GetPlayers())

	// Set bets Turn
	bets = map[string]int{
		"Player 1": 300,
		"Player 2": 350,
	}
	err = game.SetBets(bets)
	if err != nil {
		fmt.Println("Error setting bets:", err)
		return
	}
	// Show community cards after betting Turn
	showCommunityCards(game.CommunityCards)
	showPlayers(game.GetPlayers())

	// Set bets River
	bets = map[string]int{
		"Player 1": 400,
		"Player 2": 450,
	}
	err = game.SetBets(bets)
	if err != nil {
		fmt.Println("Error setting bets:", err)
		return
	}

	// Show community cards after betting River
	showCommunityCards(game.CommunityCards)
	showPlayers(game.GetPlayers())

	fmt.Println("Game Over")
	fmt.Println("Thank you for playing!")
}
