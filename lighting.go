package main

import (
	"math"
)

var SPECULAR_EXP = 4.0

func Lighting(
	normal, view, ambient []float64,
	light [][]float64,
	areflect, dreflect, sreflect []float64,
) []int {
	normal = vectorNorm(normal)
	L := light[:]
	L[0] = vectorNorm(L[0])
	view = vectorNorm(view)

	a := Ambient(ambient, areflect)
	d := Diffuse(L, dreflect, normal)
	s := Specular(L, sreflect, view, normal)

	color := make([]int, 3)
	for i := range color {
		color[i] = a[i] + d[i] + s[i]
	}
	return LimitColor(color)
}

func Ambient(alight, areflect []float64) []int {
	return vectorInt(multiplyVectors(alight, areflect))
}

func Diffuse(light [][]float64, dreflect, normal []float64) []int {
	normT := vectorDot(normal, light[0])
	if normT <= 0 {
		return []int{0, 0, 0}
	}
	kd := multiplyVectors(light[1], dreflect)
	return vectorInt(scaleVector(kd, normT))
}

func Specular(light [][]float64, sreflect, view, normal []float64) []int {
	normT := vectorDot(normal, light[0])
	if normT <= 0 {
		return []int{0, 0, 0}
	}
	normT *= 2
	R := make([]float64, len(normal))
	for i := range R {
		R[i] = normT * normal[i] - light[0][i]
	}
	highlight := math.Pow(vectorDot(R, view), SPECULAR_EXP)
	P := multiplyVectors(light[1], sreflect)
	return vectorInt(scaleVector(P, highlight))
}

func LimitColor(color []int) []int {
	res := color[:]
	for i, x := range res {
		if x > 255 {
			res[i] = 255
		}
	}
	return res
}
