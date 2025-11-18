package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите размер массива n > 2: ")
	fmt.Scan(&n)
	arr := make([]int, n)

	fmt.Print("Вводите элементы массива n раз: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Printf("Горный массив arr: %d\n", arr)
	left, right := 0, n-1

	for left < right {
		mid := (left + right) / 2

		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	fmt.Printf("Индекс пика: %d\n", right)
	fmt.Printf("Значение пика: %d\n", arr[right])
}
