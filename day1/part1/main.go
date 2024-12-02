package part1

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
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

	var distance int
	for i := len(left); i > 0; i-- {
		fmt.Println(left)
		fmt.Println(right)

		var smallestLeft, smallestRight int
		smallestLeft, left = popSmallestNumber(left)
		smallestRight, right = popSmallestNumber(right)
		fmt.Println("smallest left is ", smallestLeft)
		fmt.Println("smallest right is ", smallestRight)

		var difference int
		if smallestLeft > smallestRight {
			difference = smallestLeft - smallestRight
		} else {
			difference = smallestRight - smallestLeft
		}
		fmt.Println("difference is ", difference)

		distance += difference
		fmt.Println("distance increased to ", distance)
		fmt.Println("")
	}

	fmt.Println("Total distance is ", distance)
}

func popSmallestNumber(arr []int) (int, []int) {
	smallest := arr[0]
	smallestIndex := 0
	for i, v := range arr {
		if v < smallest {
			smallest = v
			smallestIndex = i
		}
	}
	return smallest, append(arr[:smallestIndex], arr[smallestIndex+1:]...)
}
