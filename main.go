package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите до какого числа N ищете простые числа: ")
	fmt.Scan(&n)
	size := (n - 1) / 2
	prime := make([]bool, size+1)

	for i := 0; i <= size; i++ {
		prime[i] = true
	}

	for i := 1; i <= size; i++ {
		for j := 1; i+j+2*i*j <= size; j++ {
			prime[i+j+2*i*j] = false
		}
	}

	fmt.Print("Простые числа: 2 ")
	for i := 1; i <= size; i++ {
		if prime[i] {
			fmt.Print(i*2+1, " ")
		}
	}
}
