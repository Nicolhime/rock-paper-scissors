package main

import (
	"fmt"
)

func main() {
	var player string
	fmt.Println("Welcome to Rock Paper Scissors!")
	fmt.Println("Please input either 'rock', 'paper', or 'scissors'")
	fmt.Scan(&player)                        //stores player's input
	playerHand := MapHand(player)            //converts player's input to HandThrown type
	computer := computer()                   //stores computer's input from "computer" function
	fmt.Println(fight(playerHand, computer)) //sends both inputs to the "fight" function and prints the final result of the game
}

func computer() HandThrown {
	return RandomHand()
}

func fight(player HandThrown, computer HandThrown) Winner {
	defer fmt.Println("Player chose:", player, ",", "Computer chose:", computer)

	switch {
	case player == computer:
		return Tie
	case player == Rock && computer == Paper:
		return Computer
	case player == Paper && computer == Scissors:
		return Computer
	case player == Scissors && computer == Rock:
		return Computer
	default:
		return Player
	}
}
