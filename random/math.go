package random

import "math/rand/v2"

func RandomInt() int {
	return rand.Int()
}

func RandomIntn(min int, max int) int {
	return rand.IntN(max-min) + min
}

func RandomFloat() float64 {
	return rand.Float64()
}

func RandomFloatn(min int, max int) float64 {
	return rand.Float64()*float64(max-min) + float64(min)
}
