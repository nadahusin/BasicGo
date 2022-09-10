package main

import (
	"fmt"

	"github.com/nadahusin/BasicGo/data"
	"github.com/nadahusin/BasicGo/interfaces"
)

type Row struct {
	DeretBilangan int
}

type Cube struct {
	Sisi float64
}

var deret = data.Row{DeretBilangan: 40}

func main() {
	fmt.Println("- TASK 1 -")
	fmt.Printf("%.2f\n", Round(4.37))
	fmt.Printf("%.2f\n", Round(4.32))
	fmt.Printf("%.2f\n", Round(4.324))
	fmt.Println()

	fmt.Println("- TASK 2 -")
	deret.Prima()
	fmt.Println()
	deret.Ganjil()
	fmt.Println()
	deret.Genap()
	fmt.Println()
	deret.Fibonacci()
	fmt.Println()

	fmt.Println()

	fmt.Println("- TASK 3 -")
	fmt.Println(interfaces.BangunRuang.Luas())
	fmt.Println(interfaces.BangunRuang.Keliling())
	fmt.Println(interfaces.BangunRuang.Volume())

}

func Round(value float64) float64 {
	return float64(int64(value*10+0.5)) / 10
}
