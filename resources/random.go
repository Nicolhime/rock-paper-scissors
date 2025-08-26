package rps

import (
	"math/rand"
)

func Computer() string { //called by main, generates computer's input

	rps := []string{"rock", "paper", "scissors"}

	C := (rps[rand.Intn(len(rps))])
	return C
}
