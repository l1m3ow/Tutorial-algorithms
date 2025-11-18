package main

import "fmt"

func main() {
	var n, per, k int
	fmt.Print("Введите размер массива n: ")
	fmt.Scan(&n)
	fmt.Print("Введите значение элемента, которое ищем per: ")
	fmt.Scan(&per)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	left, right := 0, n-1
	k = 0

	for left <= right {
		k++
		mid := (left + right) / 2

		if arr[mid] == per {
			fmt.Printf("Количество проверок k = %d \n", k)
			return
		} else if arr[mid] < per {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
