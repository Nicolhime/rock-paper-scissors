package main

import (
	"fmt"
	random "rps/resources"
)

func main() {
	var i string
	fmt.Println("Welcome to Rock Paper Scissors!")
	fmt.Println("Please enter either 'rock', 'paper', or 'scissors'.")
	//User input for rock, paper, or scissors was taken from here temporarily while i figure out the best location for it

	fmt.Println("value stored:", i) //temp confirmation value was stored

	fmt.Println("Generating Computer's Answer")//temp informing function is about to be called
	fmt.Println("Computer's Answer:", random.Computer()) //temp confirmation variable recieved

}
