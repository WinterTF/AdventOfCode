package main

import(
  "fmt"
  "os"
  "io"
)

func SolvePartOne(){

  fmt.Println(ReadInput())
  os.Exit(0)
  
}

func ReadInput() str{
  fileContent, error := ioutil.ReadFile("input.txt")

  if error != nil{
    fmt.Println("Error reading the input file")
    os.exit(1)
  }

  return fileContent
}
