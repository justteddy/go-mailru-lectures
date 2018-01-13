package main

import (
	"fmt"
	"reflect"
)

type memoizeFunction func(int) int

func fibonacci(n int) int {
	if n < 3 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func memoize(fn memoizeFunction) memoizeFunction {
	history := make(map[string]map[int]int)

	return func(n int) int {
		if res, ok := history[reflect.TypeOf(fn).String()]; ok {
			if val, ok := res[n]; ok {
				fmt.Println("reading from history...")
				return val
			}
		}

		val := fn(n)
		history[reflect.TypeOf(fn).String()] = map[int]int{n: val}
		return val
	}
}

func main() {
	// fibonacci test
	memoFb := memoize(fibonacci)
	fmt.Println("Fibonacci(15) =", memoFb(15))
	fmt.Println("Fibonacci(15) =", memoFb(15))
	fmt.Println("Fibonacci(15) =", memoFb(15))
}
