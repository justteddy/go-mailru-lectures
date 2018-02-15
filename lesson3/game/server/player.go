package main

import "net"

type Player struct {
	name     string
	location Locatable
	input    chan string
	conn     net.Conn
}

// NewPlayer creates a new player with name
func NewPlayer(name string, conn net.Conn) *Player {
	return &Player{
		name:  name,
		input: make(chan string),
		conn:  conn,
	}
}

func (p *Player) GetOutput() chan string {
	return p.input
}
