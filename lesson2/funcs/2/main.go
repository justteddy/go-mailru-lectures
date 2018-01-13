package main

import (
	"fmt"
	"reflect"
)

func showMeTheType(i interface{}) string {
	return reflect.TypeOf(i).String()
}

func main() {
	x := 2
	fmt.Println(showMeTheType(x))
}
