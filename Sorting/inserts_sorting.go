package main

import (
	"fmt"
)

func Inserts(arr []int) {
	fmt.Println("\n---Сортировка вставками---")
	fmt.Printf("Изначальный массив: %d\n", arr)

	size := len(arr)

	for i := 1; i < size; i++ {
		per := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > per {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = per
		fmt.Printf("\nПосле прохода номер %d по массиву\nMaссив : %d \n", i, arr)
	}
}
