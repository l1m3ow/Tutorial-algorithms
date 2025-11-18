package main

import "fmt"

func main() {
	var n int             // объявляем переменную n размер массива
	fmt.Scan(&n)          // передаём значение n через консоль
	arr := make([]int, n) // создаём массив размером n

	for i := 1; i <= n; i++ {
		arr[i-1] = i
	}

	fmt.Print(arr)
}
