package main

import (
	"fmt"
	"math"
)

func vectorSubtract(u, v []float64) (res []float64) {
	if len(u) != len(v) {
		return
	}
	res = make([]float64, len(u))
	for i := range u {
		res[i] = u[i] - v[i]
	}
	return
}

func vectorDot(u, v []float64) float64 {
	output := 0.0
	for i := range u {
		output += u[i] * v[i]
	}
	return output
}

func multiplyVectors(u, v []float64) []float64 {
	if len(u) != len(v) {
		panic(
			fmt.Sprintf(
				"Vectors %v and %v are not of the same length. Cannot multiply.",
				u, v,
			),
		)
	}
	output := make([]float64, len(u))
	for i := range u {
		output[i] = u[i] * v[i]
	}
	return output
}

func scaleVector(v []float64, a float64) []float64 {
	u := v[:]
	for i := range u {
		u[i] *= a
	}
	return u
}

func vectorInt(v []float64) []int {
	u := make([]int, len(v))
	for i, x := range v {
		u[i] = int(x)
	}
	return u
}

func vectorNorm(v []float64) []float64 {
	m := magnitude(v)
	norm := make([]float64, len(v))
	for i, x := range v {
		norm[i] = x / m
	}
	return norm
}

func magnitude(v []float64) float64 {
	vSqrd := make([]float64, len(v))
	for i, x := range v {
		vSqrd[i] = math.Pow(x, 2)
	}
	return math.Pow(sum(vSqrd), 0.5)
}

func sum(v []float64) (output float64) {
	for _, x := range v {
		output += x
	}
	return
}
