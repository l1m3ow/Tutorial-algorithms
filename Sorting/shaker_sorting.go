package main

import (
	"fmt"
	"time"
)

func Shaker(arr []int) {
	fmt.Println("\n---Шейкерная сортировка---")
	start := time.Now()

	size := len(arr)
	left := 0
	right := size - 1
	step := 1

	for left < right {
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}

		right--
		step++

		for i := right; i > left; i-- {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}

		left++
		step++
	}

	end := time.Since(start)
	fmt.Printf("Время выполнения в наносекундах: %d ns\n", end.Nanoseconds())
}
