package main

// Person is player obj
type Person struct {
	bag         []string
	isBagWeared bool
	place       Placable
}

func (p *Person) addToBag(thing string) {
	p.bag = append(p.bag, thing)
}
