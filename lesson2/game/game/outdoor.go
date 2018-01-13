package main

import "fmt"

// Outdoor is a goal location
type Outdoor struct {
	Place
}

func (o *Outdoor) string() string {
	return "улица"
}

func (o *Outdoor) look() string {
	return fmt.Sprintf("ты на улице. можно пройти - %s", accessableWays(o))
}

func (o *Outdoor) oncome() string {
	if !hallway.isDoorOpened {
		return "дверь закрыта"
	}

	player.place = o
	return "на улице весна. можно пройти - домой"
}
