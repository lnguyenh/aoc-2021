package utils

func SumIntMap(intMap map[int]int) int {
	sum := 0
	for _, num := range intMap {
		sum += num
	}
	return sum
}
