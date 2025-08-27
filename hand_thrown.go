package main

import "math/rand"

type HandThrown uint8

const (
	Rock HandThrown = iota
	Paper
	Scissors
)

var handToString = map[HandThrown]string{
	Rock:     "rock",
	Paper:    "paper",
	Scissors: "scissors",
}

func MapHand(s string) HandThrown {
	for k, v := range handToString {
		if v == s {
			return k
		}
	}
	return HandThrown(255) // invalid value
}

func (h HandThrown) String() string {
	return handToString[h]
}

func RandomHand() HandThrown {
	return HandThrown(rand.Intn(3))
}
