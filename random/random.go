package rockpaperscissors

import (
	"fmt"
	"math/rand"
)

func Rng() {

	r := []string{
		"rock",
		"paper",
		"scissors",
	}

	fmt.Println("random output:", r[rand.Intn(len(r))])

}
