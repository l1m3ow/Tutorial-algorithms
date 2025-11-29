package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Gnome(arr []int) {
	start := time.Now()

	i := 1

	for i < len(arr) {
		if i > 0 && arr[i-1] > arr[i] {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			i--
		} else {
			i++
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Время в наносекундах: %d ns\n", elapsed.Nanoseconds())
}

func main() {
	var size int
	fmt.Print("Введите размер массива: ")
	fmt.Scan(&size)

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(101)
	}

	Gnome(arr)
}
