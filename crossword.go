package main

import "fmt"

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

}
