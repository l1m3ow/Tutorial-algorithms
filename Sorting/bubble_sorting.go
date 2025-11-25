package main

import (
	"fmt"
)

func Bubble(arr []int) {
	fmt.Println("\n---Пузырьковая сортировка---")

	size := len(arr)

	for j := 0; j < size-1; j++ {
		swap := false
		fmt.Printf("\nПроход по массиву номер %d\n", j+1)
		fmt.Printf("Массив : %d \n", arr)

		for k := 0; k < size-1-j; k++ {
			if arr[k] > arr[k+1] {
				fmt.Printf("Обмен: [%d]%d ↔ [%d]%d\t", k, arr[k], k+1, arr[k+1])
				arr[k], arr[k+1] = arr[k+1], arr[k]
				swap = true
				fmt.Printf("Массив теперь : %d \n", arr)
			}
		}

		if !swap {
			break
		}
	}
}
