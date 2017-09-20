package main

import (
  "fmt"
  "log"
)

const MAX_SIZE = 9 

func printGrid(grid [][]string ){
  
  for _,row := range grid {
    for _, char := range row {
      fmt.Printf(" %v", char)
    }
    fmt.Printf("\n")
  }
}



func main(){

  grid := [][]string{ []string{"p", "a", "s", "s", "w", "0", "r", "d", "1"},
                      []string{"0", "0", "0", "0", "0", "0", "0", "0", "0"},
                      []string{"0", "0", "0", "0", "0", "0", "0", "0", "0"},
                      []string{"0", "0", "0", "0", "0", "0", "0", "0", "0"},
                      []string{"0", "0", "0", "0", "0", "0", "0", "0", "0"},
                      []string{"0", "0", "0", "0", "0", "0", "0", "0", "0"},
                      []string{"0", "0", "0", "0", "0", "0", "0", "0", "0"},
                      []string{"0", "0", "0", "0", "0", "0", "0", "0", "0"},
                      []string{"0", "0", "0", "0", "0", "0", "0", "0", "0"},
  }

  printGrid(grid)

  var searchTerm string
  fmt.Printf("Enter a word to search, or ! to exit: ")
  _, err := fmt.Scanln(&searchTerm)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(searchTerm)

}
