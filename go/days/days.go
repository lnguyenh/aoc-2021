package days

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/days/day01"
	"github.com/lnguyenh/aoc-2021/days/day02"
	"github.com/lnguyenh/aoc-2021/days/day03"
	"github.com/lnguyenh/aoc-2021/days/day04"
)

type dayFunctionType func(string)

func Run(day string, path string) {
	dayFunctions := map[string]dayFunctionType{
		"01": day01.Run,
		"02": day02.Run,
		"03": day03.Run,
		"04": day04.Run,
	}
	dayFunction, found := dayFunctions[day]
	if found {
		dayFunction(path)
	} else {
		fmt.Printf("No code ready for day %s\n", day)
	}
	return
}
