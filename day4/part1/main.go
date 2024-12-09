package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//mainFile, err := os.Open("day4/test.txt")
	//peekFile, err := os.Open("day4/test.txt")
	mainFile, err := os.Open("day4/input.txt")
	peekFile, err := os.Open("day4/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer mainFile.Close()
	defer peekFile.Close()

	mainReader := bufio.NewReader(mainFile)
	peekReader := bufio.NewReader(peekFile)

	var found int
	var linesToCheck []string

	for {
		currentLine, err := mainReader.ReadString('\n')
		if linesToCheck != nil {
			peekLine, _ := peekReader.ReadString('\n')
			// pop the first element
			linesToCheck = linesToCheck[1:]
			// add the next line
			linesToCheck = append(linesToCheck, peekLine)
		} else {
			// initialise with the current line and the next 3 lines
			for i := 0; i < 4; i++ {
				peekLine, readErr := peekReader.ReadString('\n')
				if readErr != nil {
					break
				}
				linesToCheck = append(linesToCheck, peekLine)
			}
		}

		for pos, cell := range currentLine {
			if cell == 'X' || cell == 'S' {
				if checkHorizontally(pos, currentLine) {
					//fmt.Println("Found XMAS while checking Horizontally starting from X")
					found++
				}
				if checkVertically(pos, linesToCheck) {
					//fmt.Println("Found XMAS while checking Vertically starting from X")
					found++
				}
				if checkDiagonallyOne(pos, linesToCheck) {
					//fmt.Println("Found XMAS while checking Diagonally starting from X")
					found++
				}
				if checkDiagonallyTwo(pos, linesToCheck) {
					//fmt.Println("Found XMAS while checking Diagonally starting from X")
					found++
				}
			}
		}
		if err != nil {
			break
		}
	}
	fmt.Println("Found", found, "XMAS")
}

func checkHorizontally(pos int, line string) bool {
	if pos+4 > len(line) {
		//fmt.Println("Not enough characters left")
		return false
	}

	// check if substring from pos to end of line contains XMAS
	if line[pos:(pos+4)] == "XMAS" || line[pos:(pos+4)] == "SAMX" {
		return true
	}

	return false
}

func checkVertically(pos int, followingLines []string) bool {
	if len(followingLines) != 4 {
		//fmt.Println("Not enough lines to check")
		return false
	}

	var possibleXmas string

	// iterate from line[pos] in the same column in the following lines to check if they contain XMAS
	for i := 0; i < 4; i++ {
		if followingLines[i] == "" {
			return false
		}
		possibleXmas += string(followingLines[i][pos])
	}

	if possibleXmas == "XMAS" || possibleXmas == "SAMX" {
		return true
	}

	return false
}

func checkDiagonallyOne(pos int, followingLines []string) bool {
	if len(followingLines) != 4 {
		//fmt.Println("Not enough lines to check")
		return false
	}

	var possibleXmas string

	//checkTopLeftToBottomRight()
	//checkBottomRightToTopLeft()
	for i := 0; i < 4; i++ {
		if followingLines[i] == "" {
			return false
		}
		if pos+i >= len(followingLines[i]) {
			return false
		}
		possibleXmas += string(followingLines[i][pos+i])
	}

	if possibleXmas == "XMAS" || possibleXmas == "SAMX" {
		return true
	}

	return false
}

func checkDiagonallyTwo(pos int, followingLines []string) bool {
	if len(followingLines) != 4 {
		//fmt.Println("Not enough lines to check")
		return false
	}

	var possibleXmas string

	//checkTopRightToBottomLeft()
	//checkBottomLeftToTopRight()
	for i := 0; i < 4; i++ {
		if followingLines[i] == "" {
			return false
		}
		if pos-i < 0 {
			return false
		}
		possibleXmas += string(followingLines[i][pos-i])
	}

	if possibleXmas == "XMAS" || possibleXmas == "SAMX" {
		return true
	}

	return false
}
