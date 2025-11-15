package main

import (
	"fmt"
	"time"
)

func Shella(arr []int) {
	fmt.Println("\n---Сортировка алгоритмом Шелла---")
	start := time.Now()

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

	end := time.Since(start)
	fmt.Printf("Время выполнения в наносекундах: %d ns\n", end.Nanoseconds())
}
