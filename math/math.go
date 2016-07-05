package math

func Average(arr []float64) float64 {
	var sum float64 = 0
	for _, value := range arr {
		sum += value
	}
	return sum / float64(len(arr))
}
