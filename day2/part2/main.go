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

		fmt.Println("checking report", report)
		if validateReport(report) || validateReducedReport(report) {
			numSafeReports++
		}

		if err != nil {
			break
		}
	}

	fmt.Println("Number of safe reports found ", numSafeReports)
}

func validateReducedReport(report []int) bool {
	fmt.Println("checking if reduced versions of report", report, "are safe")
	var reducedReport []int
	for i := 0; i < len(report); i++ {
		reducedReport = removeIndexFromReport(i, report)
		fmt.Println("reduced report", reducedReport)

		if validateReport(reducedReport) {
			fmt.Println("reduced report", reducedReport, "is safe")
			return true
		}
	}
	return false
}

func removeIndexFromReport(index int, report []int) []int {
	fmt.Println("removing index", index, "from report", report)
	newReport := make([]int, len(report)-1)
	copy(newReport, report[:index])
	copy(newReport[index:], report[index+1:])
	return newReport
}

func validateReport(report []int) bool {
	var increasing, decreasing bool

	for i := 0; i < len(report); i++ {
		var difference int

		if i == len(report)-1 {
			break
		}

		current := report[i]
		next := report[i+1]

		if increasing {
			if current >= next {
				return false
			} else {
				difference = next - current
			}
		} else if decreasing {
			if current <= next {
				return false
			} else {
				difference = current - next
			}
		} else {
			// first level

			if current < next {
				increasing = true
				difference = next - current
			} else if current > next {
				decreasing = true
				difference = current - next
			} else if current == next {
				return false
			}
		}

		if difference < 1 || difference > 3 {
			return false
		}
	}

	fmt.Println("report", report, "is safe")
	return true
}
