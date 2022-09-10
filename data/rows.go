package data

import (
	"fmt"
)

type Row struct {
	DeretBilangan int
}

func (x *Row) Prima() []int {
	var value int
	var result []int
	for i := 1; i <= x.DeretBilangan; i++ {
		value = 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				value = value + 1

				result = append(result, i)
			}
		}

		if value == 2 {
			fmt.Print(i, " ")
		}
	}
	return result
}

func (x *Row) Ganjil() []int {
	var result []int
	for i := 0; i <= x.DeretBilangan; i++ {
		if i%2 != 0 {
			fmt.Print(i, " ")
			result = append(result, i)
		}
	}
	return result
}

func (x *Row) Genap() []int {
	var result []int
	for i := 0; i <= x.DeretBilangan; i++ {
		if i%2 == 0 && i > 0 {
			fmt.Print(i, " ")
			result = append(result, i)

		}
	}
	return result

}

func (x *Row) Fibonacci() []int {
	var result []int
	var fib1 int = 0 // sebagai nilai pertama deret bilangan fibonacci.
	var fib2 = 1     // sebagai bilangan kedua deret bilangan fibonacci.
	var sum = 1      // untuk mencetak result dari kedua bilangan sebelumnya (fib1 dan fib2) yang telah dijumlahkan

	fmt.Print(fib1, " ")
	for sum <= x.DeretBilangan {
		fmt.Print(sum, " ")
		sum = fib1 + fib2
		fib1 = fib2
		fib2 = sum

		result = append(result, sum)
	}

	return result

}
