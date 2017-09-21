package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Directions to search in when matching
const DIRECTIONS = 8

// PrintGrid takes the 2D slice representing the crossword
// it formats the grid in to a nice presentation
func printGrid(grid [][]string) {

	for _, row := range grid {
		for _, char := range row {
			fmt.Printf(" %v", char)
		}
		fmt.Printf("\n")
	}
}

// Search takes the word to search for and the 2D slice representing the crossword
// it returns the start index row, column, and direction of the word
// if the word isn't found, returns -1 for these 3 values
// direction works as if you mapped 1 to NE on a compass, and circled clockwise
func search(word string, grid [][]string) (frow, fcol, fdir int) {

	frow, fcol, fdir = -1, -1, -1

	for ridx, row := range grid {
		for cidx := range row {
			for dir := 1; dir <= DIRECTIONS; dir++ {
				if match(word, grid, ridx, cidx, dir) {
					frow, fcol, fdir = ridx+1, cidx+1, dir
					return
				}
			}
		}
	}
	return
}

// SearchAll takes the list of words provided and searches all of them automatically
func searchAll(words []string, grid [][]string) {

	for _, word := range words {
		row, col, dir := search(word, grid)
		printSearch(word, row, col, dir)
	}

}

// PrintSearch nicely formats the search results
func printSearch(searchTerm string, row, col, dir int) {

	fmt.Printf("Word: %v AT (%v,%v,%v) (r,c,d)\n", searchTerm, row, col, dir)

}

// OutOfBounds is a helper function for match
// It takes the indexes and dimensions
// Then checks to make sure the search does not go off the grid
func outOfBounds(ridx, cidx, rowlength, collength int) (out bool) {

	if cidx >= collength || cidx < 0 || ridx >= rowlength || ridx < 0 {
		out = true
	}

	return

}

// MatchChar determines if two characters match
func matchChar(char, gridChar string) (match bool) {

	if char == gridChar {
		match = true
	}

	return

}

// Helper function for match that makes checks bounds and character matching
func matchChecks(char, word string, grid [][]string, ridx, cidx, dir, pos, rowlength, collength int) (matches bool) {

	if outOfBounds(ridx, cidx, rowlength, collength) {
		return
	}
	if !matchChar(char, grid[ridx][cidx]) {
		return
	}
	if matchChar(char, grid[ridx][cidx]) && pos+1 == len(word) {
		matches = true
	}

	return
}

// Match searches the crossword grid in direction starting at
// ridx, cidx, and returns true if word is found
func match(word string, grid [][]string, ridx, cidx, dir int) (found bool) {

	rowlength, collength := len(grid), len(grid[0])
	switch dir {
	//NE
	case 1:
		for pos, chrune := range word {
			char := string(chrune)
			if !matchChecks(char, word, grid, ridx, cidx, dir, pos, rowlength, collength) {
				break
			} else {
				found = true
			}
			ridx--
			cidx++
		}
	//E
	case 2:
		for pos, chrune := range word {
			char := string(chrune)
			if !matchChecks(char, word, grid, ridx, cidx, dir, pos, rowlength, collength) {
				break
			} else {
				found = true
			}
			cidx++
		}
	//SE
	case 3:
		for pos, chrune := range word {
			char := string(chrune)
			if !matchChecks(char, word, grid, ridx, cidx, dir, pos, rowlength, collength) {
				break
			} else {
				found = true
			}
			ridx++
			cidx++
		}
	//S
	case 4:
		for pos, chrune := range word {
			char := string(chrune)
			if !matchChecks(char, word, grid, ridx, cidx, dir, pos, rowlength, collength) {
				break
			} else {
				found = true
			}
			ridx++
		}
	//SW
	case 5:
		for pos, chrune := range word {
			char := string(chrune)
			if !matchChecks(char, word, grid, ridx, cidx, dir, pos, rowlength, collength) {
				break
			} else {
				found = true
			}
			ridx++
			cidx--
		}
	//W
	case 6:
		for pos, chrune := range word {
			char := string(chrune)
			if !matchChecks(char, word, grid, ridx, cidx, dir, pos, rowlength, collength) {
				break
			} else {
				found = true
			}
			cidx--
		}
	//NW
	case 7:
		for pos, chrune := range word {
			char := string(chrune)
			if !matchChecks(char, word, grid, ridx, cidx, dir, pos, rowlength, collength) {
				break
			} else {
				found = true
			}
			ridx--
			cidx--
		}
	//N
	case 8:
		for pos, chrune := range word {
			char := string(chrune)
			if !matchChecks(char, word, grid, ridx, cidx, dir, pos, rowlength, collength) {
				break
			} else {
				found = true
			}
			ridx--
		}
	}

	return

}

// Main program logic, searches a crossword puzzle for words
// Also determines if a wordlist should be used, and searches those words automatically
func main() {

	args := os.Args

	if len(args) == 1 {
		fmt.Println("Usage: crossword <gridfile> <word bank (optional)>")
		os.Exit(1)
	}

	gridFile, err := os.Open(args[1])
	if err != nil {
		panic(err)
	}

	defer gridFile.Close()

	//make grid
	var lines []string
	gridScanner := bufio.NewScanner(gridFile)
	for gridScanner.Scan() {
		lines = append(lines, gridScanner.Text())
	}
	var grid [][]string
	for _, line := range lines {
		var splitLine []string
		for _, chrune := range line {
			char := string(chrune)
			splitLine = append(splitLine, char)
		}
		grid = append(grid, splitLine)
	}

	//make wordlist
	var words []string
	if len(args) == 3 {

		wordFile, err := os.Open(args[2])
		if err != nil {
			panic(err)
		}
		wordScanner := bufio.NewScanner(wordFile)
		for wordScanner.Scan() {
			words = append(words, wordScanner.Text())
		}

		defer wordFile.Close()

	}

	printGrid(grid)

	fmt.Println(words)

	if len(args) == 2 {
		var searchTerm string
		fmt.Printf("Enter a word to search, or ! to exit: ")
		for searchTerm != "!" {
			_, err := fmt.Scanln(&searchTerm)
			if err != nil {
				log.Fatal(err)
			}
			if searchTerm == "!" {
				os.Exit(0)
			}
			row, col, dir := search(searchTerm, grid)
			printSearch(searchTerm, row, col, dir)

			fmt.Printf("Enter a word to search, or ! to exit: ")
		}
	} else {
		searchAll(words, grid)
	}

}
