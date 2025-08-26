package fight

import (
	"fmt"
	"math/rand"
)

func Fight(p1 string, p2 string) string {
	
	switch expression {
	case condition:
		
	}
	fmt.Scan(&p1) //user gets to choose!
	Computer() //computer gets to choose!

	return
}

func Computer() string { 
	rps := []string{"rock", "paper", "scissors"}
	return (rps[rand.Intn(len(rps))]) //randomly chooses an option from the slice above and returns i
}
