package main

import (
	"fmt"

	"math/rand"
)

func main() {
	var player string
	fmt.Println("Welcome to Rock Paper Scissors!")
	fmt.Println("Please input either 'rock', 'paper', or 'scissors'")
	fmt.Scan(&player)                    //stores player's input
	computer := computer()               //stores computer's input from "computer" function
	fmt.Println(fight(player, computer)) //sends both inputs to the "fight" function and prints the final result of the game
}

func computer() string {
	rps := []string{"rock", "paper", "scissors"}
	return (rps[rand.Intn(len(rps))]) //randomly chooses an option from the slice above and returns it
}

func fight(player string, computer string) string {
	var output string

	switch {
	case player == computer:
		fmt.Println("Player chose:", player, ",", "Computer chose:", computer) //if both inputs are the same, output tie
		output = "Its a Tie!"
	case player == "rock" && computer == "paper":
		fmt.Println("Player chose:", player, ",", "Computer chose:", computer) //if computer beats player, output loss
		output = "Computer Wins!"
	case player == "paper" && computer == "scissors":
		fmt.Println("Player chose:", player, ",", "Computer chose:", computer)
		output = "Computer Wins!"
	case player == "scissors" && computer == "rock":
		fmt.Println("Player chose:", player, ",", "Computer chose:", computer)
		output = "Computer Wins!"
	case player == "rock" && computer == "scissors":
		fmt.Println("Player chose:", player, ",", "Computer chose:", computer) //if player beats computer, output win
		output = "Player Wins!"
	case player == "paper" && computer == "rock":
		fmt.Println("Player chose:", player, ",", "Computer chose:", computer)
		output = "Player Wins!"
	case player == "scissors" && computer == "paper":
		fmt.Println("Player chose:", player, ",", "Computer chose:", computer)
		output = "Player Wins!"
	}
  
	return output
}
