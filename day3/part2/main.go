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
	var instructionsEnabled = true

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		sumMult += processLine(line, &instructionsEnabled)
		//fmt.Println("Sum of all multiplications found so far: ", sumMult)
		if err != nil {
			break
		}
	}

	fmt.Println("Sum of all multiplications found: ", sumMult)
}

const START = 0
const MulStateM = 1
const MulStateU = 2
const MulStateL = 3
const MulStateOpen = 4
const MulStateComma = 5
const MulStateClose = 6
const DoOrDontStateD = 7
const DoOrDontStateO = 8
const DoStateOpen = 9
const DoStateClose = 10
const DontStateN = 11
const DontStateSingleQuote = 12
const DontStateT = 13
const DontStateOpen = 14
const DontStateClose = 15

func processLine(line string, instructionsEnabled *bool) int {
	var num1, num2 string
	var sumMult int

	state := START

	for i := 0; i < len(line); i++ {
		char := line[i]
		//fmt.Println("processing char: ", string(char))

		switch state {
		case START:
			if char == 'm' {
				state = MulStateM
			} else if char == 'd' {
				state = DoOrDontStateD
			} else {
				state = START
			}
		case MulStateM:
			if char == 'u' {
				state = MulStateU
			} else {
				state = START
			}
		case MulStateU:
			if char == 'l' {
				state = MulStateL
			} else {
				state = START
			}
		case MulStateL:
			if char == '(' {
				state = MulStateOpen
				num1 = ""
			} else {
				state = START
			}
		case MulStateOpen:
			if char >= '0' && char <= '9' {
				num1 += string(char)
			} else if char == ',' {
				state = MulStateComma
				num2 = ""
			} else {
				state = START
			}
		case MulStateComma:
			if char >= '0' && char <= '9' {
				num2 += string(char)
			} else if char == ')' {
				state = MulStateClose
				if num1 != "" && num2 != "" {
					if n1, err := strconv.Atoi(num1); err == nil {
						if n2, err := strconv.Atoi(num2); err == nil {
							if *instructionsEnabled {
								//fmt.Printf("Found valid instruction: mul(%s,%s)\n", num1, num2)
								//fmt.Println(n1, "*", n2, " = ", n1*n2)
								//fmt.Println("Adding", n1*n2, "to", sumMult, "equals", sumMult+n1*n2)
								sumMult += n1 * n2
							}
						}
					}
				}
				state = START
			} else {
				state = START
			}
		case DoOrDontStateD:
			if char == 'o' {
				state = DoOrDontStateO
			} else {
				state = START
			}
		case DoOrDontStateO:
			if char == '(' {
				state = DoStateOpen
			} else if char == 'n' {
				state = DontStateN
			} else {
				state = START
			}
		case DoStateOpen:
			if char == ')' {
				state = DoStateClose
				*instructionsEnabled = true
				//fmt.Println("do() detected, enabling instructions")
			}
			state = START
		case DontStateN:
			if char == '\'' {
				state = DontStateSingleQuote
			} else {
				state = START
			}
		case DontStateSingleQuote:
			if char == 't' {
				state = DontStateT
			} else {
				state = START
			}
		case DontStateT:
			if char == '(' {
				state = DontStateOpen
			} else {
				state = START
			}
		case DontStateOpen:
			if char == ')' {
				state = DontStateClose
				*instructionsEnabled = false
				//fmt.Println("don't() detected, disabling instructions")
			}
			state = START
		default:
			state = START
		}
	}
	return sumMult
}
