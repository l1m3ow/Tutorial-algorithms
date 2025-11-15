package main

import (
	"fmt"
	"time"
)

func Choice(arr []int) {
	fmt.Println("\n---Сортировка выбором---")
	start := time.Now()

	size := len(arr)

	for i := 0; i < size; i++ {
		min_index := i

		for j := i + 1; j < size; j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}

		arr[i], arr[min_index] = arr[min_index], arr[i]
	}

	end := time.Since(start)
	fmt.Printf("Время выполнения в наносекундах: %d ns\n", end.Nanoseconds())
}
