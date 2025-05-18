package main

import (
	"fmt"
	"math"
)

func main() {
	xValues := []float64{0.2, 0.3, 0.38, 0.43, 0.57}

	fmt.Println("y = (sin^3(x) + cos^3(x)) * ln(x):")
	for _, x := range xValues {
		y := (math.Pow(math.Sin(x), 3) + math.Pow(math.Cos(x), 3)) * math.Log(x)
		fmt.Println("При x =", x, "y =", y)
	}
}