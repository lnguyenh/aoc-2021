package days

import (
	"fmt"
	"github.com/lnguyenh/aoc-2021/days/day01"
	"github.com/lnguyenh/aoc-2021/days/day02"
	"github.com/lnguyenh/aoc-2021/days/day03"
	"github.com/lnguyenh/aoc-2021/days/day04"
	"github.com/lnguyenh/aoc-2021/days/day05"
	"github.com/lnguyenh/aoc-2021/days/day06"
	"github.com/lnguyenh/aoc-2021/days/day07"
	"github.com/lnguyenh/aoc-2021/days/day08"
	"github.com/lnguyenh/aoc-2021/days/day09"
	"time"
)

type dayFunctionType func(string)

func Run(day string, path string) {
	startTime := time.Now()

	dayFunctions := map[string]dayFunctionType{
		"01": day01.Run,
		"02": day02.Run,
		"03": day03.Run,
		"04": day04.Run,
		"05": day05.Run,
		"06": day06.Run,
		"07": day07.Run,
		"08": day08.Run,
		"09": day09.Run,
	}
	dayFunction, found := dayFunctions[day]
	if found {
		dayFunction(path)
	} else {
		fmt.Printf("No code ready for day %s\n", day)
	}

	fmt.Println(time.Since(startTime))
	return
}
