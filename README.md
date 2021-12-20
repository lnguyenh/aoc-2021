# Advent of code 2021

## About
OK, let's do this another year! This time, I will use `Advent of Code` to learn Go. Let's... go!

https://adventofcode.com/2021

## Setup
Input files are expected to be located in the `inputs` folder under the correct folder for the day, and to have the extension `.txt`. Example `inputs/01/test.txt`

## Running the Go solutions

- First you need to install Go and have it available in your path so that it is possible to call the go command in a terminal
- The, to get the solution for `day01` and for the input file  `inputs/01/input.txt`, run the following:

```
cd go
go run cmd/aoc.go input 01
```
Alternatively, to solve today's problem with the corresponding input `input.txt`:
```
go run cmd/aoc.go input
```

## Understanding the solutions
A good starting point is to check first the `solution.go` file for a specific day, look at the  `Run()` function for each day and follow the code from there.

## Adding a new Go solution
- Create a new input directory in the `inputs` directory, typically with a `test.txt` and a `input.txt` file as provided on https://adventofcode.com/
- Create a new directory in the `go/days` directory, using a previous day as a template
- Add a key/value pair in the `dayFunctions` map in the file `go/days/days.go`

Voila, start coding in `go/days/[your new day]/solution.go`!