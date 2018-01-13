package main

// Placable interface for game places
type Placable interface {
	oncome() string
	Wearer
	Enver
	Looker
	Stringer
	Stuffer
}

type Wearer interface {
	wears() *[]string
}
type Enver interface {
	envs() *[]string
}
type Looker interface {
	look() string
}

type Stringer interface {
	string() string
}

type Stuffer interface {
	stuffs() *[]string
}
