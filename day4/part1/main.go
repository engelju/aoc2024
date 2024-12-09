package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("day4/test.txt")
	//f, err := os.Open("day4/input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var found int
	var unreadBuffer bytes.Buffer

	var linesToCheck []string

	r := bufio.NewReader(f)
	for {
		var currentLine string

		// Read from the buffer first if there are unread lines
		if unreadBuffer.Len() > 0 {
			currentLine, err = unreadBuffer.ReadString('\n')
		} else {
			currentLine, err = r.ReadString('\n')
		}

		for pos, cell := range currentLine {
			if cell == 'X' {
				//fmt.Println("Found an X")
				linesToCheck = nil

				if checkLeftToRight(pos, currentLine) {
					fmt.Println("Found XMAS while checking Left to Right")
					found++
				}

				if checkRightToLeft(pos, currentLine) {
					fmt.Println("Found XMAS while checking Right to Left")
					found++
				}

				linesToCheck = append(linesToCheck, currentLine)
				for i := 0; i < 3; i++ {
					followingLine, readErr := r.ReadString('\n')
					linesToCheck = append(linesToCheck, followingLine)

					// Write the peeked lines into the unreadBuffer
					unreadBuffer.WriteString(followingLine)

					if readErr != nil {
						break
					}
				}

				if checkTopToBottom(pos, linesToCheck) {
					fmt.Println("Found XMAS while checking Top to Bottom")
					found++
				}

				//checkBottomToTop()
				//checkTopLeftToBottomRight()
				//checkBottomRightToTopLeft()
				//checkTopRightToBottomLeft()
				//checkBottomLeftToTopRight()
			}
		}

		if err != nil {
			break
		}
	}
	fmt.Println("Found", found, "XMAS")
}

func checkLeftToRight(pos int, line string) bool {

	if pos+4 > len(line) {
		//fmt.Println("Not enough characters left")
		return false
	}

	// check if substring from pos to end of line contains XMAS
	if line[pos:(pos+4)] == "XMAS" {
		return true
	}

	return false
}

func checkRightToLeft(pos int, line string) bool {

	if pos-4 < 0 {
		return false
	}

	// check if substring from pos to start of line contains XMAS
	substr := line[pos-3 : pos+1]
	if substr == "SAMX" {
		return true
	}

	return false
}

func checkTopToBottom(pos int, followingLines []string) bool {

	if len(followingLines) != 4 {
		//fmt.Println("Not enough lines to check")
		return false
	}

	var possibleXmas string

	// iterate from line[pos] in the same column in the following lines to check if they contain XMAS
	for i := 0; i < 4; i++ {
		fmt.Println("Checking", string(followingLines[i][pos]))
		possibleXmas += string(followingLines[i][pos])
	}

	if possibleXmas == "XMAS" {
		return true
	}

	return false
}
