package main

import (
	"fmt"
	"time"
)

func Inserts(arr []int) {
	fmt.Println("\n---Сортировка вставками---")
	start := time.Now()

	size := len(arr)

	for i := 1; i < size; i++ {
		per := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > per {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = per
	}

	end := time.Since(start)
	fmt.Printf("Время выполнения в наносекундах: %d ns\n", end.Nanoseconds())
}
