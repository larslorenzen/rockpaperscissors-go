package main

import (
	"testing"
)

func TestBeats(t *testing.T) {
	rockPlayer := Player{ROCK}
	if rockPlayer.beats(Player{ROCK}) {
		t.Error("Should not beat same hand")
	}
	if rockPlayer.beats(Player{SCISSORS}) {
		t.Error("Rock does not beat scissors")
	}
	if (!rockPlayer.beats(Player{PAPER})) {
		t.Error("Rock must beat paper")
	}
}
