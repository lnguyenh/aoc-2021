package days

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/days/day01"
	"github.com/lnguyenh/aoc-2021/days/day02"
)

type dayFunctionType func(string)

func Run(day string, path string) {
	functions := map[string]dayFunctionType{
		"01": day01.Run,
		"02": day02.Run,
	}
	dayFunction, found := functions[day]
	if found {
		dayFunction(path)
	} else {
		fmt.Printf("No code ready for day %s\n", day)
	}
	return
}
