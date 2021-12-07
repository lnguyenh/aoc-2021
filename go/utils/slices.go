package utils

func SumSlice(measurement []int) int {
	sum := 0
	for _, value := range measurement {
		sum += value
	}
	return sum
}

func MaxSlice(slice []int) int {
	max := 0
	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	return max
}

func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IntSliceToSet(list []int) []int {
	var set []int
	for _, value := range list {
		if !IntInSlice(value, set) {
			set = append(set, value)
		}
	}
	return set
}
