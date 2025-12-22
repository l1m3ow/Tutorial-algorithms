package main

import "fmt"

func Sum(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + Sum(n/10)
}

func main() {
	var n int
	fmt.Print("Введите число n: ")
	fmt.Scan(&n)
	ot := Sum(n)
	fmt.Printf("Сумма чисел числа %d = %d", n, ot)
}
