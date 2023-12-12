package main

import (
  "fmt"
  "os"
  "github.com/TimFWinter/AdventOfCode/2023/day01"
)

func main()  {
  if len(os.Args) != 2 {
    fmt.Println("Invalid arguments")
    os.Exit(1)
  }

  puzzle := os.Args[1]

  switch puzzle{

  case "1.1":
    day01.SolvePartOne()

  default:
    fmt.Println("Invalid puzzle selected")

  }
}
