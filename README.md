# aoc2022

Welcome to my solutions to Advent of Code 2022! This year, I have (again) chosen to go with Go. Some resons why:

- The problems usually present really good test cases. Go has nice testing out-of-the-box.
- Some cases really lend themselves to "table structured" testing. This is easily represented in Go.
- Go gives us a lot with just the standard library.

## General structure

- Each day has its own package
- Each day has `Run()` function, called from the main function. This outputs the final values.
- Each day has a private `parse()` function, since every day's puzzle will require different parsing.
- After that, the implementations follow
- [testify's](https://github.com/stretchr/testify) `assert` package is used to keep the test code cleaner.

## How to

### Run

From the repository root, run
> go run main.go

### Run tests
To run tests for all packages, from the repository root, run
> go test ./...

If you want to run tests for only `day1`, run
> go test ./day1 