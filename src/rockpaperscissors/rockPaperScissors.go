package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Hand int
}

func (h *Player) beats(other Player) bool {
	return h.Hand == (other.Hand+1)%3
}

const (
	ROCK     = iota
	SCISSORS
	PAPER
)

var (
	Wins   = 0
	Draws  = 0
	Losses = 0
)

func play(random *rand.Rand) {
	rock := Player{ROCK}
	randomHand := Player{random.Intn(3)}
	if randomHand == rock {
		Draws++
	} else if randomHand.beats(rock) {
		Wins++
	} else {
		Losses++
	}
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		play(random)
	}
	fmt.Println("Wins = ", Wins, " Draws = ", Draws, " Losses = ", Losses)
}
