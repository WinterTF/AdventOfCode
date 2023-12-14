package day01

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func SolvePartOne() {

	puzzleInput := strings.Splits(ReadInput(), `\n`)

	os.Exit(0)

}

func ReadInput() string {
	fileContent, error := ioutil.ReadFile("inputs/input1.txt")

	if error != nil {
		fmt.Println("Error reading the input file")
		os.Exit(1)
	}

	return string(fileContent)
}

func ExtractDigits(line string) list[string] {

}
