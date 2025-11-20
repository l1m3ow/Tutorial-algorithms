package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите до какого числа N ищете простые числа: ")
	fmt.Scan(&n)
	prime := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		prime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if prime[i] {
			for j := i * i; j <= n; j += i {
				prime[j] = false
			}
		}
	}

	fmt.Print("Простые числа: ")
	for i := 2; i <= n; i++ {
		if prime[i] {
			fmt.Print(i, " ")
		}
	}
}
