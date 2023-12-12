package main

import (
  "fmt"
  "os"
)

func main()  {
  if len(os.args) != {
    fmt.Println("Invalid arguments")
    os.Exit(1)
  }

  puzzle := os.Args[1]

  switch puzzle{

  case "1.1":
    solveDay1(1)

  default:
    ftm.Println("Invalid puzzle selected")

  }
}