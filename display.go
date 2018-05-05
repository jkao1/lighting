// display contains useful functions for the screen
package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"os/exec"
)

const XRES = 500
const YRES = 500
const PPMFilename = "pic.ppm"

// NewScreen creates a new screen of size XRES by YRES. It returns the new
// screen.
func NewScreen() (screen [][][]int) {
	screen = make([][][]int, YRES)
	for i := range screen {
		screen[i] = make([][]int, XRES)
		for j := range screen[i] {
			screen[i][j] = []int{255, 255, 255}
		}
	}
	return
}

// NewZBuffer creates a new z-buffer of size XRES by YRES. It returns the new
// z-buffer.
func NewZBuffer() (zbuffer [][]float64) {
	zbuffer = make([][]float64, YRES)
	for i := range zbuffer {
		zbuffer[i] = make([]float64, XRES)
		for j := range zbuffer[i] {
			zbuffer[i][j] = float64(math.MinInt64)
		}
	}
	return
}

// DisplayScreen uses XQuartz's "display" command to display a PPM.
func DisplayScreen(screen [][][]int) {
	WriteScreenToPPM(screen)
	_, err := exec.Command("display", PPMFilename).Output()
	if err != nil {
		panic(err)
	}
}

// ClearScreen clears a screen.
func ClearScreen(screen [][][]int) {
	for i, _ := range screen {
		screen[i] = make([][]int, XRES)

		for j, _ := range screen[i] {
			screen[i][j] = []int{255, 255, 255}
		}
	}
}

// WriteScreenToExtension writes a screen to a filename.
func WriteScreenToExtension(screen [][][]int, filename string) {
	WriteScreenToPPM(screen)
	_, err := exec.Command("convert", PPMFilename, filename).Output()
	if err != nil {
		panic(err)
	}
}

// WriteScreenToPPM takes a screen as an argument and writes it to a PPM file.
func WriteScreenToPPM(screen [][][]int) {
	file, err := os.OpenFile(PPMFilename, os.O_CREATE | os.O_WRONLY, 0644)
	if (err != nil) {
		panic(err)
	}

	defer file.Close()

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("P3 %d %d 255\n", XRES, YRES))
	for i := 0; i < YRES; i++ {
		for j := 0; j < XRES; j++ {
			rgb := screen[i][j]
			buffer.WriteString(fmt.Sprintf("%d %d %d ", uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2])))
		}
	}

	file.WriteString(buffer.String())
}
