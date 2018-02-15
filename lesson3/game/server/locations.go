package main

type Locatable interface {
	getLocationName() string
}

type Hallway struct {
}

func (h *Hallway) getLocationName() string {
	return "hallway"
}

type Kitchen struct {
}

func (k *Kitchen) getLocationName() string {
	return "kitchen"
}
