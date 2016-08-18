package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	// Should only provide a single file for now, not empty
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		os.Exit(1)
	}
	if len(args) == 0 {
		fmt.Println("No file provided")
		os.Exit(1)
	}

	// Safe to try receiving json file
	fileName := args[0]
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Could not retrieve file: ", err)
		os.Exit(1)
	}

	var list []cell
	json.Unmarshal(file, &list)
	if err != nil {
		fmt.Println("Could not Unmarshal json: ", err)
		os.Exit(1)
	}

	fmt.Println("Loaded cells:")
	fmt.Println("------------")
	for _, c := range list {
		fmt.Printf("%v x %v\n", c.X, c.Y)
	}
}

type cell struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}
type cellMap map[cell]bool

var (
	curBoard = make(cellMap)
	newBoard = make(cellMap)
)

func (m cellMap) AddVec(v cell) {
	m[v] = true
}

func setup() {

}

func update() {

}

func draw() {

}
