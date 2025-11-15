package main

import (
	"fmt"
	"time"
)

func Bubble(arr []int) {
	fmt.Println("\n---Пузырьковая сортировка---")
	start := time.Now()

	size := len(arr)

	for j := 0; j < size-1; j++ {
		swap := false

		for k := 0; k < size-1-j; k++ {
			if arr[k] > arr[k+1] {
				arr[k], arr[k+1] = arr[k+1], arr[k]
				swap = true
			}
		}

		if !swap {
			break
		}
	}
	end := time.Since(start)
	fmt.Printf("Время выполнения в наносекундах: %d ns\n", end.Nanoseconds())
}
