package main

import (
	"math/rand"
	"strings"
)

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

func (h HandThrown) String() string {
	return handToString[h]
}

func MapHand(s string) HandThrown {
	s = strings.ToLower(s)
	for k, v := range handToString {
		if v == s {
			return k
		}
	}
	return HandThrown(255) // invalid value
}

func RandomHand() HandThrown {
	return HandThrown(rand.Intn(3))
}
