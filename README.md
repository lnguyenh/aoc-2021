# Advent of code 2021

## About
OK, let's try another year!

https://adventofcode.com/2021

## Setup
Input files are expected to be located in the "inputs" folder under the correct folder for the day. Example "inputs/01/test.txt"

## Running the Go solutions

In order to get the solution for day 01 and for the input file  "inputs/01/input.txt", run the following:

```
cd go
go run cmd/aoc.go 01 input.txt
```

## Adding a new Go solution
- Create a new input folder with empty files
- Create a new "day" folder
- Add a key/value pair in the dayFunctions map in "days.go"
