package internal

func Average(nums []float64) float64 {
	sum := 0.0

	for _, n := range nums {
		sum += n
	}

	avg := sum / float64(len(nums))

	return avg
}
