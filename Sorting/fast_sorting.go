package main

import (
	"fmt"
	"time"
)

func Fast(arr []int) {
	fmt.Println("\n---Быстрая сортировка---")
	start := time.Now()

	size := len(arr)

	var sort func(int, int)

	sort = func(low, high int) {

		if low < high {
			per := arr[high]
			i := low - 1

			for j := low; j < high; j++ {
				if arr[j] <= per {
					i++
					arr[i], arr[j] = arr[j], arr[i]
				}
			}

			arr[i+1], arr[high] = arr[high], arr[i+1]
			perIndex := i + 1
			sort(low, perIndex-1)
			sort(perIndex+1, high)
		}
	}

	sort(0, size-1)

	end := time.Since(start)
	fmt.Printf("Время выполнения в наносекундах: %d ns\n", end.Nanoseconds())
}
