package internal

func Median(numbers []float64) float64 {
	var median float64
	if len(numbers)%2 == 1 {
		median = numbers[len(numbers)/2]
	} else {
		median = (numbers[len(numbers)/2-1] + numbers[len(numbers)/2]) / 2
	}
	return median
}
