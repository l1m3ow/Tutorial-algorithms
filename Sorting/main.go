package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var size int
	fmt.Print("Введите размер массива: ")
	fmt.Scan(&size)

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(101)
	}

	Bubble(arr)
}
