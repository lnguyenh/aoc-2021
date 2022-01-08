package day24

func doType1(i, oldZ, dy int) int {
	return 26*oldZ + (i + dy)
}

func doType26(i, oldZ, dx, dy int) int {
	x := oldZ%26 + dx

	if x == i {
		return oldZ / 26
	} else {
		return oldZ/26*26 + (i + dy) // replaces last term
	}
}

func doProgram(number []int) int {
	z := 0
	z = doType1(number[0], 0, 12) // #1 always < 26  (z = i1 +12)
	z = doType1(number[1], z, 9)  // #2 *26 [+10, + 18] --
	z = doType1(number[2], z, 8)  // #3 *26 [+9, +17]  --

	z = doType26(number[3], z, -8, 3) // #4 --

	z = doType1(number[4], z, 0)  // #5 *26 [+1, +9] --
	z = doType1(number[5], z, 11) // #6 *26 [+12, +20] --
	z = doType1(number[6], z, 10) // #7 *26 [+11, +19] --

	z = doType26(number[7], z, -11, 13) // #8 --

	z = doType1(number[8], z, 3) // #9 *26 [+4, +12] --

	z = doType26(number[9], z, -1, 10)  // #10 --
	z = doType26(number[10], z, -8, 10) // #11 --
	z = doType26(number[11], z, -5, 14) // #12 --
	z = doType26(number[12], z, -16, 6) // #13 --
	z = doType26(number[13], z, -6, 5)  // #14

	return z
}
