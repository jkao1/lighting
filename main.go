package main

func main() {
	view := []float64{0, 0, 1}
	ambient := []float64{50, 50, 50}
	light := [][]float64{
		[]float64{0.5, 0.75, 1},
		[]float64{236, 120, 59},
	}
	areflect := []float64{.1, .1, .1}
	dreflect := []float64{0.5, 0.5, 0.5}
	sreflect := []float64{0.5, 0.5, 0.5}

	screen := NewScreen()
	zbuffer := NewZBuffer()
	transform := make([][]float64, 0)
	edges := make([][]float64, 4)

	ParseFile("script", transform, edges, screen, zbuffer, view, ambient, light, areflect, dreflect, sreflect)
}
