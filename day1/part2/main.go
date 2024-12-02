package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//f, err := os.Open("day1/test.txt")
	f, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var left, right []int
	r := bufio.NewReader(f)
	for {
		var firstNum, secondNum int
		if _, err := fmt.Fscan(r, &firstNum, &secondNum); err != nil {
			break
		}
		left = append(left, firstNum)
		right = append(right, secondNum)
	}

	var numUsages, totalSimilarityScore int
	for i := 0; i < len(left); i++ {
		numUsages = findUsages(left[i], right)
		fmt.Println("Number of usages for ", left[i], " is ", numUsages)

		simScore := numUsages * left[i]
		totalSimilarityScore += simScore
	}

	fmt.Println("Total similarity score is ", totalSimilarityScore)
}

func findUsages(num int, arr []int) int {
	var count int
	for _, n := range arr {
		if num == n {
			count++
		}
	}
	return count
}
