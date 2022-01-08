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
	"github.com/lnguyenh/aoc-2021/days/day10"
	"github.com/lnguyenh/aoc-2021/days/day11"
	"github.com/lnguyenh/aoc-2021/days/day12"
	"github.com/lnguyenh/aoc-2021/days/day13"
	"github.com/lnguyenh/aoc-2021/days/day14"
	"github.com/lnguyenh/aoc-2021/days/day15"
	"github.com/lnguyenh/aoc-2021/days/day16"
	"github.com/lnguyenh/aoc-2021/days/day17"
	"github.com/lnguyenh/aoc-2021/days/day18"
	"github.com/lnguyenh/aoc-2021/days/day19"
	"github.com/lnguyenh/aoc-2021/days/day20"
	"github.com/lnguyenh/aoc-2021/days/day21"
	"github.com/lnguyenh/aoc-2021/days/day22"
	"github.com/lnguyenh/aoc-2021/days/day23"
	"github.com/lnguyenh/aoc-2021/days/day24"
	"github.com/lnguyenh/aoc-2021/days/day25"
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
		"10": day10.Run,
		"11": day11.Run,
		"12": day12.Run,
		"13": day13.Run,
		"14": day14.Run,
		"15": day15.Run,
		"16": day16.Run,
		"17": day17.Run,
		"18": day18.Run,
		"19": day19.Run,
		"20": day20.Run,
		"21": day21.Run,
		"22": day22.Run,
		"23": day23.Run,
		"24": day24.Run,
		"25": day25.Run,
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
