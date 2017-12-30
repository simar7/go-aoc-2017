package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func sToI(s string) int {
	// check if it's a valid number
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return int(num)
	}
	log.Fatal(err)
	return 0
}

func part1() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var maxVal int
	for _, line := range lines {
		//fmt.Println(line)
		lineMax := math.MinInt64
		lineMin := math.MaxInt64

		// break into individual line number words
		lineNums := strings.Fields(line)

		// iterate over each line word
		for val := range lineNums {
			// check if it's a valid number
			if num, err := strconv.ParseInt(string(lineNums[val]), 10, 64); err == nil {
				//fmt.Println(num)
				if int(num) >= lineMax {
					lineMax = int(num)
				}
				if int(num) <= lineMin {
					lineMin = int(num)
				}
			}

		}
		//fmt.Println("line: ", k, "max: ", lineMax, "min: ", lineMin)
		valDiff := lineMax - lineMin
		maxVal = maxVal + valDiff
	}

	fmt.Println("[part1]", maxVal)
}

func part2() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var maxVal int
	for _, line := range lines {
		//fmt.Println(line)
		lineNum1 := 1
		lineNum2 := 1
		pairFound := false

		// break into individual line number words
		lineNums := strings.Fields(line)

		// iterate over each line word
		// dirty O(N^2)
		for val := range lineNums {
			if pairFound == false {
				for otherVal := range lineNums {
					// both numbers are equal, skip
					// TODO: can lines have same values?
					if lineNums[val] == lineNums[otherVal] {
						continue
					}

					// both numbers divide each other
					if sToI(string(lineNums[val]))%sToI(string(lineNums[otherVal])) == 0 || sToI(string(lineNums[otherVal]))%sToI(string(lineNums[val])) == 0 {
						lineNum1 = sToI(string(lineNums[val]))
						lineNum2 = sToI(string(lineNums[otherVal]))
						fmt.Println("they divide cleanly: ", sToI(string(lineNums[val])), sToI(string(lineNums[otherVal])))
						if lineNum1 > lineNum2 {
							maxVal = maxVal + (lineNum1 / lineNum2)
						} else {
							maxVal = maxVal + (lineNum2 / lineNum1)
						}
						pairFound = true
						break
					}
				}
			}
		}
	}

	fmt.Println("[part2]", maxVal)
}

func main() {
	part1()
	part2()
}
