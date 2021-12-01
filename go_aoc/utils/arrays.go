package utils

func SumArray(measurement []int) int {
	sum := 0
	for _, value := range measurement {
		sum += value
	}
	return sum
}
