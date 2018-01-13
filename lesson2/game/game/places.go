package main

import (
	"log"
	"strings"
)

// Place is a basic type for game rooms
type Place struct {
	stuff []string
	wear  []string
	env   []string
}

func (p *Place) stuffs() *[]string {
	return &p.stuff
}

func (p *Place) envs() *[]string {
	return &p.env
}

func (p *Place) wears() *[]string {
	return &p.wear
}

func accessableWays(place Placable) string {
	if _, ok := worldmap[place]; !ok {
		log.Fatal("неизвестное место")
	}

	ways := []string{}
	for _, variant := range worldmap[place] {
		ways = append(ways, variant.string())
	}

	return strings.Join(ways, ", ")
}

func getFromCurrentPlace(thing string, container string) bool {

	var copied *[]string
	var slice []string

	switch container {
	case "stuff":
		copied = player.place.stuffs()
		slice = *player.place.stuffs()
	case "wear":
		copied = player.place.wears()
		slice = *player.place.wears()
	}

	for i, st := range slice {
		if st == thing {
			*copied = append(slice[:i], slice[i+1:]...)
			return true
		}
	}

	return false
}
