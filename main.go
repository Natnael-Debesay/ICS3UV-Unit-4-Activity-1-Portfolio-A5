// Author: Natanel Debesay
// Version: 1.0.0
// Date 2025-12-09
// Fileoverview: This program prints the ASCII value between two prompted numbers between 32 and 126.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variable declaration
	var startingCharacterASCIIString string
	var startingCharacterASCIINumber int
	var endingCharacterASCIIString string
	var endingCharacterASCIINumber int

	var reader = bufio.NewReader(os.Stdin)

	// input numbers
	fmt.Print("Please enter a number larger than 32, and less than 126: ")
	startingCharacterASCIIString, _ = reader.ReadString('\n')
	startingCharacterASCIIString = strings.TrimSpace(startingCharacterASCIIString)
	startingCharacterASCIINumber, _ = strconv.Atoi(startingCharacterASCIIString)

	fmt.Print("Please enter a number larger than 32, and less than 126: ")
	endingCharacterASCIIString, _ = reader.ReadString('\n')
	endingCharacterASCIIString = strings.TrimSpace(endingCharacterASCIIString)
	endingCharacterASCIINumber, _ = strconv.Atoi(endingCharacterASCIIString)

	// print numbers
	for counter := startingCharacterASCIINumber; counter <= endingCharacterASCIINumber; counter ++ {
		startingCharacterASCIIString = string(rune(startingCharacterASCIINumber))
		endingCharacterASCIIString = string(rune(endingCharacterASCIINumber))
		fmt.Printf("%s\n", counter)
	}

	fmt.Println("\nDone.")
}
