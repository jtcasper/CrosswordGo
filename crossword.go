package main

import (
  "fmt"
  "log"
  "os"
)

// Directions to search in when matching
const DIRECTIONS = 8

// PrintGrid takes the 2D slice representing the crossword
// it formats the grid in to a nice presentation
func printGrid(grid [][]string ){
  
  for _,row := range grid {
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
func search(word string, grid [][]string)(frow, fcol, fdir int){

  frow, fcol, fdir = -1, -1, -1
  
  for ridx, row := range grid {
    for cidx, _ := range row {
      for dir := 1; dir <= DIRECTIONS; dir++ {
        if match( word, grid, ridx, cidx, dir ) {
          frow, fcol, fdir = ridx + 1, cidx + 1, dir
          return
        }
      }
    }
  }
  return
}

// Match searches the crossword grid in direction starting at
// ridx, cidx, and returns true if word is found
func match( word string, grid [][]string, ridx, cidx, dir int )(found bool){

  rowlength, collength := len(grid), len(grid[0])
  switch dir {
    //NE
    case 1:
      for pos,chrune := range word {
        char := string(chrune)
//      fmt.Printf("r %v, c %v\n", ridx, cidx)
        if cidx >= collength || ridx < 0 {
          break
        }
        if grid[ridx][cidx] != char{
          break
        }
        if grid[ridx][cidx] == char && pos + 1 == len(word){
          found = true
//        fmt.Printf("Found %v\n", char)
        }
        ridx--
        cidx++
      }
    //E
    case 2:
      for pos,chrune := range word {
        char := string(chrune)
        if cidx >= collength {
          break
        }
        if grid[ridx][cidx] != char{
          break
        }
        if grid[ridx][cidx] == char && pos + 1 == len(word){
          found = true
        }
        cidx++
      }
    //SE
    case 3:
      for pos,chrune := range word {
        char := string(chrune)
        if cidx >= collength || ridx >= rowlength{
          break
        }
        if grid[ridx][cidx] != char{
          break
        }
        if grid[ridx][cidx] == char && pos + 1 == len(word){
          found = true
        }
        ridx++
        cidx++
      }
    //S
    case 4:
      for pos,chrune := range word {
        char := string(chrune)
        if ridx >= rowlength{
          break
        }
        if grid[ridx][cidx] != char{
          break
        }
        if grid[ridx][cidx] == char && pos + 1 == len(word){
          found = true
        }
        ridx++
      }
    //SW
    case 5:
      for pos,chrune := range word {
        char := string(chrune)
        if cidx < 0 || ridx >= rowlength{
          break
        }
        if grid[ridx][cidx] != char{
          break
        }
        if grid[ridx][cidx] == char && pos + 1 == len(word){
          found = true
        }
        ridx++
        cidx--
      }
    //W
    case 6:
      for pos,chrune := range word {
        char := string(chrune)
        if cidx < 0 {
          break
        }
        if grid[ridx][cidx] != char{
          break
        }
        if grid[ridx][cidx] == char && pos + 1 == len(word){
          found = true
        }
        cidx--
      }
    //NW
    case 7:
      for pos,chrune := range word {
        char := string(chrune)
        if cidx < 0 || ridx < 0 {
          break
        }
        if grid[ridx][cidx] != char{
          break
        }
        if grid[ridx][cidx] == char && pos + 1 == len(word){
          found = true
        }
        ridx--
        cidx--
      }
    //N
    case 8:
      for pos,chrune := range word {
        char := string(chrune)
        if ridx < 0 {
          break
        }
        if grid[ridx][cidx] != char{
          break
        }
        if grid[ridx][cidx] == char && pos + 1 == len(word){
          found = true
        }
        ridx--
      }
  }

  return

}

// Main program logic, searches a crossword puzzle for words
func main(){

  grid := [][]string{ []string{"p", "a", "s", "s", "w", "0", "r", "d", "1"},
                      []string{"p", "0", "0", "0", "0", "0", "0", "0", "0"},
                      []string{"0", "0", "0", "0", "0", "0", "0", "0", "0"},
                      []string{"0", "0", "0", "0", "0", "0", "0", "0", "0"},
                      []string{"0", "0", "0", "0", "0", "0", "0", "0", "0"},
                      []string{"0", "0", "0", "0", "0", "0", "0", "0", "0"},
                      []string{"0", "0", "0", "0", "0", "0", "0", "0", "0"},
                      []string{"0", "0", "0", "0", "0", "0", "0", "0", "0"},
                      []string{"b", "0", "0", "0", "0", "0", "0", "0", "0"},
                      []string{"a", "0", "0", "0", "0", "0", "0", "0", "0"},
  }

  printGrid(grid)

  var searchTerm string
  fmt.Printf("Enter a word to search, or ! to exit: ")
  for searchTerm != "!" {
    _, err := fmt.Scanln(&searchTerm)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(searchTerm)
    if searchTerm == "!" {
      os.Exit(0)
    }

    fmt.Println(search(searchTerm, grid))

    fmt.Printf("Enter a word to search, or ! to exit: ")
  }

}
