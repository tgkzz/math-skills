package internal

import "math"

func Variance(numbers []float64, average float64) float64 {
	var sumOfSquares float64
	for _, number := range numbers {
		sumOfSquares += math.Pow(number-average, 2)
	}
	return sumOfSquares / float64(len(numbers)-1)
}
