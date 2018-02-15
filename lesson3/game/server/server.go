package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var players = map[string]*Player{}
var incoming = make(chan *Player)
var leaving = make(chan *Player)
var broadcast = make(chan string)

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

func broadcastMonitor() {
	for {
		select {
		case player := <-incoming:
			fmt.Fprintf(os.Stdout, "%s is connected\n", player.name)
			players[player.name] = player
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
		case len(parsedCmd) == 2 && parsedCmd[0] == "сказать":
			broadcast <- parsedCmd[1]
		case len(parsedCmd) == 3 && parsedCmd[0] == "сказать_игроку":
			target, ok := players[parsedCmd[1]]
			if !ok {
				fmt.Fprintf(c, "Player %s is not existed\n", parsedCmd[1])
				continue
			}
			target.input <- "Player " + player.name + " says: " + parsedCmd[2]
		}
	}

	leaving <- player
}
