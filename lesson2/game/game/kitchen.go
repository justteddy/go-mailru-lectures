package main

import (
	"fmt"
	"strings"
)

// Kitchen is a start location
type Kitchen struct {
	Place
}

func (k *Kitchen) string() string {
	return "кухня"
}

func (k *Kitchen) look() string {
	var stuff string
	var isbag string
	if len(k.stuff) == 0 {
		stuff = "ничего нет"
	} else {
		stuff = strings.Join(k.stuff, ", ")
	}

	if !player.isBagWeared {
		isbag = " собрать рюкзак и"
	}

	return fmt.Sprintf("ты находишься на кухне, на столе %s, надо%s идти в универ. можно пройти - %s", stuff, isbag, accessableWays(k))
}

func (k *Kitchen) oncome() string {
	player.place = k
	return fmt.Sprintf("кухня, ничего интересного. можно пройти - %s", accessableWays(k))
}
