package main

import (
	"fmt"

	"github.com/elferherrera/dswithgo/section2/parser"
)

func main() {
	newData := parser.NewData(10)

	for _, val := range newData {
		parser.PrettyPrint(val)
	}

	parser.SaveToFile(newData, "data.json")

	fmt.Println("Reading file")
	readAgain, _ := parser.LoadFile("data.json")

	for _, val := range readAgain {
		parser.PrettyPrint(val)
	}
}
