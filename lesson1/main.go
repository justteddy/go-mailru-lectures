package main

import (
	"sort"
	"strconv"
)

func main() {

}

func ReturnInt() int {
	var result int
	result = 1
	return result
}

func ReturnFloat() float32 {
	var result float32
	result = 1.1
	return result
}

func ReturnIntArray() [3]int {
	result := [3]int{1, 3, 4}
	return result
}

func ReturnIntSlice() []int {
	result := []int{1, 2, 3}
	return result
}

func IntSliceToString(intSlice []int) string {
	var result string
	for _, val := range intSlice {
		result += strconv.Itoa(val)
	}

	return result
}

func MergeSlices(floatSlice []float32, intSlice []int32) []int {
	result := make([]int, 0, len(floatSlice)+len(intSlice))

	for _, val := range floatSlice {
		result = append(result, int(val))
	}

	for _, val := range intSlice {
		result = append(result, int(val))
	}

	return result
}

func GetMapValuesSortedByKey(input map[int]string) []string {
	var keys []int
	var result []string
	for k := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		result = append(result, input[k])
	}

	return result
}
