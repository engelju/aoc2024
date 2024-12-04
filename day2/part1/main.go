package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//f, err := os.Open("day2/test.txt")
	f, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var numSafeReports int

	r := bufio.NewReader(f)
	for {
		// each line is a report, consisting of numbers representing levels that are separated by spaces

		line, err := r.ReadString('\n')

		var report []int
		for _, num := range strings.Split(line, " ") {
			var n int
			fmt.Sscan(num, &n)
			report = append(report, n)
		}

		if checkIfReportIsSafe(report) {
			fmt.Println("Report is safe")
			numSafeReports++
		} else {
			fmt.Println("Report is not safe")
		}

		if err != nil {
			break
		}
	}

	fmt.Println("Number of safe reports found ", numSafeReports)
}

/**
 * A report only counts as safe if both of the following are true:
 * 	- The levels are either all increasing or all decreasing.
 * 	- Any two adjacent levels differ by at least one and at most three.
 */
func checkIfReportIsSafe(report []int) bool {
	var increasing, decreasing bool

	fmt.Println(report)
	for i := 0; i < len(report); i++ {
		var difference int

		if i == len(report)-1 {
			break
		}

		current := report[i]
		next := report[i+1]

		if increasing {
			if current > next {
				return false
			} else {
				difference = next - current
			}
		} else if decreasing {
			if current < next {
				return false
			} else {
				difference = current - next
			}
		} else if current < next {
			increasing = true
			difference = next - current
		} else if current > next {
			decreasing = true
			difference = current - next
		}

		if difference < 1 || difference > 3 {
			return false
		}
	}

	return true
}
