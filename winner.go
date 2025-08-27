package main

type Winner uint8

const (
	Tie Winner = iota
	Player
	Computer
)

func (w Winner) String() string {
	switch w {
	case Tie:
		return "It's a tie!"
	case Player:
		return "Player wins!"
	case Computer:
		return "Computer wins!"
	default:
		return "Unknown result"
	}
}
