package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//worldmap
var worldmap map[Placable][]Placable

// places
var kitchen Kitchen
var room Room
var outdoor Outdoor
var hallway Hallway

//thingcallback
var thingcallback map[string]func(string) string

// player
var player Person

func main() {
	initGame()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(handleCommand(scanner.Text()))
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
}

func handleCommand(command string) string {
	defer func() {
		fmt.Println("-----------------")
	}()

	parsedCmd := strings.Split(command, " ")

	switch {
	case len(parsedCmd) == 1 && parsedCmd[0] == "хелп":
		showHelp()
		return ""
	case len(parsedCmd) == 1 && parsedCmd[0] == "осмотреться":
		return look()
	case len(parsedCmd) == 2 && parsedCmd[0] == "идти":
		return move(parsedCmd[1])
	case len(parsedCmd) == 2 && parsedCmd[0] == "одеть":
		return wear(parsedCmd[1])
	case len(parsedCmd) == 2 && parsedCmd[0] == "взять":
		return put(parsedCmd[1])
	case len(parsedCmd) == 3 && parsedCmd[0] == "применить":
		return use(parsedCmd[1], parsedCmd[2])
	}

	return "неизвестная команда"
}

func initGame() {
	kitchen.stuff = []string{"чай"}

	room.stuff = []string{"ключи", "конспекты"}
	room.wear = []string{"рюкзак"}

	outdoor = Outdoor{}

	hallway.env = []string{"дверь"}

	player.place = &kitchen

	worldmap = map[Placable][]Placable{
		&kitchen: []Placable{&hallway},
		&hallway: []Placable{&kitchen, &room, &outdoor},
		&room:    []Placable{&hallway},
		&outdoor: []Placable{&hallway},
	}

	thingcallback = map[string]func(string) string{
		"ключи": func(target string) string {
			hallway.isDoorOpened = true
			return fmt.Sprintf("%s открыта", target)
		},
		"рюкзак": func(target string) string {
			player.isBagWeared = true
			return fmt.Sprintf("вы одели: %s", target)
		},
	}

	showHelp()
}

func showHelp() {
	fmt.Println("------------------------------------")
	fmt.Println("хелп - напомнить команды")
	fmt.Println("осмотреться - что происходит вокруг тебя")
	fmt.Println("взять <предмет> - взять <предмет>")
	fmt.Println("одеть <вещь> - одеть на себя <вещь>")
	fmt.Println("идти <место> - пойти в <место>")
	fmt.Println("применить <предмет> <цель> - применить <предмет> на определенную <цель>")
}
