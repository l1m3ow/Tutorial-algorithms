package main

import (
	"fmt"
	"math/rand"
)

func Shella(arr []int) {
	fmt.Println("\n---Сортировка алгоритмом Шелла---")

	size := len(arr)
	gap := size / 2

	for gap > 0 {
		for i := gap; i < size; i++ {
			per := arr[i]
			j := i

			for j >= gap && arr[j-gap] > per {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = per
		}
		gap /= 2
	}
}

func main() {
	var size int
	fmt.Print("Введите размер массива: ")
	fmt.Scan(&size)

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(101)
	}

	Shella(arr)
}
