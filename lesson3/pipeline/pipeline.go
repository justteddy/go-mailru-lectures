package main

import (
	"fmt"
)

type job func(in, out chan interface{})

// Pipe running few functions in a pipeline
func Pipe(funcs ...job) {
	ch := make(chan interface{})
	for _, step := range funcs {
		ch = func(step job, in chan interface{}) chan interface{} {
			out := make(chan interface{})
			go step(in, out)

			return out
		}(step, ch)
	}

	for res := range ch {
		fmt.Println(res)
	}

}

func main() {
	firstFunc := job(func(in, out chan interface{}) {
		for i := 0; i < 10; i++ {
			out <- i * 10
		}
		close(out)
	})

	secondFunc := job(func(in, out chan interface{}) {
		for i := range in {
			out <- i.(int) + 1
		}
		close(out)
	})

	thirdFunc := job(func(in, out chan interface{}) {
		for i := range in {
			out <- i.(int) - 1
		}
		close(out)
	})

	Pipe(firstFunc, secondFunc, thirdFunc)
}
