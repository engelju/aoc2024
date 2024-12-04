package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//f, err := os.Open("day3/test.txt")
	f, err := os.Open("day3/input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var sumMult int

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		fmt.Println(line)
		sumMult += processLine(line)
		if err != nil {
			break
		}
	}

	fmt.Println("Sum of all multiplications found: ", sumMult)
}

func processLine(line string) int {
	state := 0
	var num1, num2 string
	var sumMult int

	for i := 0; i < len(line); i++ {
		char := line[i]

		switch state {
		case 0:
			if char == 'm' {
				state = 1
			}
		case 1:
			if char == 'u' {
				state = 2
			} else {
				state = 0
			}
		case 2:
			if char == 'l' {
				state = 3
			} else {
				state = 0
			}
		case 3:
			if char == '(' {
				state = 4
				num1 = ""
			} else {
				state = 0
			}
		case 4:
			if char >= '0' && char <= '9' {
				num1 += string(char)
			} else if char == ',' {
				state = 5
				num2 = ""
			} else {
				state = 0
			}
		case 5:
			if char >= '0' && char <= '9' {
				num2 += string(char)
			} else if char == ')' {
				if num1 != "" && num2 != "" {
					if n1, err := strconv.Atoi(num1); err == nil {
						if n2, err := strconv.Atoi(num2); err == nil {
							fmt.Printf("Found valid instruction: mul(%s,%s)\n", num1, num2)
							sumMult += n1 * n2
						}
					}
				}
				state = 0
			} else {
				state = 0
			}
		}
	}
	return sumMult
}
