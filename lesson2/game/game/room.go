package main

import (
	"fmt"
	"strings"
)

// Room is a main room
type Room struct {
	Place
}

func (r *Room) string() string {
	return "комната"
}

func (r *Room) look() string {
	var stuff string
	var wear string

	if len(r.stuff) == 0 {
		stuff = "пустая комната"
	} else {
		stuff = fmt.Sprintf("на столе: %s", strings.Join(r.stuff, ", "))
	}

	if len(r.wear) == 0 {
		wear = ""
	} else {
		wear = fmt.Sprintf(", на стуле - %s", strings.Join(r.wear, ", "))
	}

	return fmt.Sprintf("%s%s. можно пройти - %s", stuff, wear, accessableWays(r))
}

func (r *Room) oncome() string {
	player.place = r
	return fmt.Sprintf("ты в своей комнате. можно пройти - %s", accessableWays(r))
}
