package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for _, i := range CreateWithSize(100) {
		fmt.Println(i)
	}
}

func Population() []int {
	population := make([]int, 100)
	for i, _ := range population {
		population[i] = rand.Intn(2)
	}
	return population
}

func CreateWithSize(size int) [][]int {
	population := make([][]int, size)
	for i, _ := range population {
		population[i] = Population()
	}
	return population
}
