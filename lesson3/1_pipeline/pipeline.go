package main

import (
	"math"
)

type job func(in, out chan interface{})

// Pipe running few functions in a pipeline
func Pipe(funcs ...job) {
	in := make(chan interface{})
	out := make(chan interface{})
	for _, pipePart := range funcs {
		go pipePart(in, out)
	}
}

func main() {
	firstFunc := job(func(in, out chan interface{}) {
		for i := 0; i < 10; i++ {
			out <- i * i
		}
	})

	secondFunc := job(func(in, out chan interface{}) {
		for i := range in {
			out <- interface{}(math.Sqrt(float64(i.(int))))
		}
	})

	Pipe(firstFunc, secondFunc)
}
