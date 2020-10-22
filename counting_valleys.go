package main

import "fmt"

func countingValleys(steps int, path string) int {
	numberOfValleys, delta := 0, 0
	for _, step := range path {
		if step == 'U' {
			delta++
			if delta == 0 {
				numberOfValleys++
			}

		} else {
			delta--
		}

	}
	return numberOfValleys

}

func main() {
	firstPath, secondPath := "DDUUUUDD", "DDUUUUUDUDDDDUDU"
	for _, path := range []string{firstPath, secondPath} {
		fmt.Println(countingValleys(len(path), path))
	}

}
