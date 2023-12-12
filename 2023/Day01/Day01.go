package day01

import(
  "fmt"
  "os"
  "io/ioutil"
)

func SolvePartOne(){

  fmt.Println(ReadInput())
  os.Exit(0)
  
}

func ReadInput() string{
  fileContent, error := ioutil.ReadFile("inputs/input1.txt")

  if error != nil{
    fmt.Println("Error reading the input file")
    os.Exit(1)
  }

  return string(fileContent)
}
