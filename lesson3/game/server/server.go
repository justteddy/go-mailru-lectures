package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

var players = map[string]*Player{}
var incoming = make(chan *Player)
var leaving = make(chan *Player)
var looking = make(chan *Player)
var broadcast = make(chan string)

var kitchen = new(Kitchen)
var hallway = new(Hallway)

var roomLoad sync.Once

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcastMonitor()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func getLocationByName(name string) (Locatable, error) {
	switch name {
	case "hallway":
		return hallway, nil
	case "kitchen":
		return kitchen, nil
	}

	return nil, errors.New("location not found")
}

func broadcastMonitor() {
	for {
		select {
		case player := <-incoming:
			fmt.Fprintf(os.Stdout, "%s is connected\n", player.name)
			players[player.name] = player
			player.location = hallway
			go func() {
				for msg := range player.input {
					fmt.Fprintf(player.conn, "%s\n", msg)
				}
			}()
		case player := <-leaving:
			delete(players, player.name)
			fmt.Fprintf(os.Stdout, "%s is disconnected\n", player.name)
		case msg := <-broadcast:
			for _, player := range players {
				player.input <- msg
			}
		case player := <-looking:
			var nearPlayers []string
			for _, gamer := range players {
				if player.name == gamer.name {
					continue
				}
				if gamer.location == player.location {
					nearPlayers = append(nearPlayers, gamer.name)
				}
			}
			fmt.Fprintf(player.conn, "You are now in %s, near - %s\n", player.location.getLocationName(), strings.Join(nearPlayers, ", "))
		default:
			continue
		}
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)

	//get player name
	fmt.Fprint(c, "Enter your name:\n")
	for input.Scan() {
		if _, ok := players[input.Text()]; ok {
			fmt.Fprintf(c, "Name %s is allready in use:\n", input.Text())
			continue
		}
		break
	}

	//create player
	player := NewPlayer(input.Text(), c)
	incoming <- player

	for input.Scan() {
		parsedCmd := strings.Split(input.Text(), " ")

		switch {
		case parsedCmd[0] == "say":
			message := strings.TrimPrefix(input.Text(), "say ")
			broadcast <- player.name + " screaming:" + message
		case parsedCmd[0] == "whisper":
			target, ok := players[parsedCmd[1]]
			if !ok {
				fmt.Fprintf(c, "Player %s is not existed\n", parsedCmd[1])
				continue
			}
			target.input <- "Player " + player.name + " whispers: " + strings.Join(parsedCmd[2:], " ")
		case parsedCmd[0] == "move":
			location, err := getLocationByName(parsedCmd[1])
			if err != nil {
				fmt.Fprintf(c, "Location %s is not existed\n", parsedCmd[1])
				continue
			}
			player.location = location
		case parsedCmd[0] == "look":
			looking <- player
		}
	}

	leaving <- player
}
