package main

import (
	"fmt"
)

func look() string {
	return player.place.look()
}

func move(place string) string {
	for _, variant := range worldmap[player.place] {
		if variant.string() == place {
			return variant.oncome()
		}
	}

	return fmt.Sprintf("нет пути в %s", place)
}

func wear(wear string) string {
	isHere := false
	for _, wr := range *player.place.wears() {
		if wr == wear {
			isHere = true
		}
	}

	if !isHere {
		return "нет такого"
	}

	if !getFromCurrentPlace(wear, "wear") {
		return "нет такого"
	}

	callback, ok := thingcallback[wear]
	if !ok {
		return "нечего одеть"
	}

	return callback(wear)
}

func put(thing string) string {
	if !player.isBagWeared {
		return "некуда класть"
	}

	if !getFromCurrentPlace(thing, "stuff") {
		return "нет такого"
	}

	player.addToBag(thing)
	return fmt.Sprintf("предмет добавлен в инвентарь: %s", thing)
}

func use(thing, target string) string {
	inBag := false
	envHere := false
	for _, stuff := range player.bag {
		if stuff == thing {
			inBag = true
			break
		}
	}

	if !inBag {
		return fmt.Sprintf("нет предмета в инвентаре - %s", thing)
	}

	for _, env := range *player.place.envs() {
		if target == env {
			envHere = true
			break
		}
	}

	if !envHere {
		return "не к чему применить"
	}

	callback, ok := thingcallback[thing]
	if !ok {
		return "не к чему применить"
	}

	return callback(target)
}
