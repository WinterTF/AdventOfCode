package main

import (
  "fmt"
  "os"
  "github.com/TimFWinter/AdventOfCode/2023/Day01"
)

func main()  {
  if len(os.args) != {
    fmt.Println("Invalid arguments")
    os.Exit(1)
  }

  puzzle := os.Args[1]

  switch puzzle{

  case "1.1":
    Day1.SolvePartOne()

  default:
    ftm.Println("Invalid puzzle selected")

  }
}
