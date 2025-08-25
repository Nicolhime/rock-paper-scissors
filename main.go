package main

import "fmt"

func main() {

	var i string
	fmt.Println("Welcome to Rock Paper Scissors!")
	fmt.Println("Please enter either 'rock', 'paper', or 'scissors'.")

	fmt.Scan(&i)
	fmt.Println("value stored:", i)
}
