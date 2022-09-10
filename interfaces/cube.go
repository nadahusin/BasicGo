package interfaces

import "math"

type Hitung2d interface {
	Luas() float64
	Keliling() float64
}

type Hitung3d interface {
	Volume() float64
}

type Hitung interface {
	Hitung2d
	Hitung3d
}

type Cube struct {
	Sisi float64
}

var BangunRuang Hitung = &Cube{4}

func (y *Cube) Luas() float64 {
	return math.Pow(y.Sisi, 2) * 6

}

func (y *Cube) Keliling() float64 {
	return y.Sisi * 12

}

func (y *Cube) Volume() float64 {
	return math.Pow(y.Sisi, 3)

}
