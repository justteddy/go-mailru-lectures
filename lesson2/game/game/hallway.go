package main

import "fmt"

// Hallway is a "crossroad" location
type Hallway struct {
	isDoorOpened bool
	Place
}

func (h *Hallway) string() string {
	return "коридор"
}

func (h *Hallway) look() string {
	return fmt.Sprintf("ты в коридоре. можно пройти - %s", accessableWays(h))
}

func (h *Hallway) oncome() string {
	player.place = h
	return fmt.Sprintf("ничего интересного. можно пройти - %s", accessableWays(h))
}
