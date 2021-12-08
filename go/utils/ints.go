package utils

func IntAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func IntIn(value int, list []int) bool {
	for _, okValue := range list {
		if value == okValue {
			return true
		}
	}
	return false
}
