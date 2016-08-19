package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var aliveChar, edgeChar, deadChar string

func main() {
	alivePtr := flag.String("alive", "X", "Representing a living cell")
	edgePtr := flag.String("edge", " ", "Representing an edge cell")
	deadPtr := flag.String("dead", " ", "Representing a dead cell")
	timePtr := flag.Int("frame-ms", 500, "Milliseconds between frames")
	flag.Parse()

	// Just get the first letter
	aliveChar = (*alivePtr)[:1]
	edgeChar = (*edgePtr)[:1]
	deadChar = (*deadPtr)[:1]
	frameDelay := time.Duration(*timePtr) * time.Millisecond

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
	// From here, all initial cells have been loaded into vertex list 'list'
	// Now to add it to the current board
	for _, v := range list {
		curBoard[cell{v.X, v.Y}] = alive
	}

	// For the first step, pass through the current board and activate edges
	for k, v := range curBoard {
		if v == alive {
			for _, edge := range k.neighbors() {
				curBoard.addEdge(edge)
			}
		}
	}

	for a := 0; a < width+2; a++ {
		fmt.Print("_")
	}
	fmt.Println()

	for {
		draw()
		curBoard.step(newBoard)
		curBoard = make(cellMap)
		for k, v := range newBoard {
			curBoard[k] = v
		}
		newBoard = make(cellMap)
		time.Sleep(frameDelay)
	}
}

const (
	width  = 30
	height = 20
)

// The key is the location itself to prevent duplicates, and the value is as
// Follows:
//	0: dead
//	1: alive
//	2: edge (dead, but prospective in the next generation)
type cellMap map[cell]cellState

type cell struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type cellState byte

const (
	dead  cellState = iota
	alive cellState = iota
	edge  cellState = iota
)

var (
	// Stores all the current cells
	curBoard = make(cellMap)
	// Stores the cells of the next generation
	newBoard = make(cellMap)
)

// If the cell is alive, that's that
func (cMap cellMap) addAlive(c cell) {
	cMap[c] = alive
}

// If the cell is an neighbor, it could also be alive. Best not to overwrite.
func (cMap cellMap) addEdge(c cell) {
	val, ok := cMap[c]
	if !ok || val == dead {
		cMap[c] = edge
	}
}

// Performs all the checking at once, may need to be split up. Fills new with
// the next generation of cells. Looks pretty messy to me, could be cleaned up
func (cMap cellMap) step(new cellMap) {
	for c, stat := range cMap {
		// Checking all non-dead cells
		if stat > 0 {
			// count the neighbors
			neighbors := c.neighbors()
			count := 0
			for _, n := range neighbors {
				if cMap[n] == alive {
					count++
				}
			}
			if count == 3 { // Cell is alive
				new.addAlive(c)
				// Add all neighbors as edges
				for _, n := range neighbors {
					new.addEdge(n)
				}
			} else if count == 2 { // The Cell doesn't change
				// Add all neighbors as edges if cell would be alive
				if stat == alive {
					new.addAlive(c)
					for _, n := range neighbors {
						new.addEdge(n)
					}
				}
				if stat == edge {
					new.addEdge(c)
				}
			}
		}
	}
}

func (c cell) neighbors() []cell {
	return []cell{
		cell{c.X + 1, c.Y + 1},
		cell{c.X + 1, c.Y - 1},
		cell{c.X - 1, c.Y - 1},
		cell{c.X - 1, c.Y + 1},
		cell{c.X + 1, c.Y},
		cell{c.X - 1, c.Y},
		cell{c.X, c.Y - 1},
		cell{c.X, c.Y + 1},
	}
}

func update() {

}

// Draws the current board
func draw() {
	// Testing the step
	var x int32
	var y int32

	fmt.Print("|")
	for i := 0; i < width; i++ {
		fmt.Print(" ")
	}
	fmt.Println("|")
	for y = 0; y < height; y++ {
		fmt.Print("|")
		for x = 0; x < width; x++ {
			stat := curBoard[cell{x, y}]
			switch {
			case stat == dead:
				fmt.Print(deadChar)
			case stat == alive:
				fmt.Print(aliveChar)
			case stat == edge:
				fmt.Print(edgeChar)
			}
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("|")
	for a := 0; a < width; a++ {
		fmt.Print("_")
	}
	fmt.Println("|")
}
