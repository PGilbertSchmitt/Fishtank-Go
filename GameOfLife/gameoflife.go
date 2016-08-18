package main

import (
	"github.com/pgilbertschmitt/fishtank/util"
)

func main() {

}

type cellMap map[util.Vertex]bool

var (
	curBoard = make(cellMap)
	newBoard = make(cellMap)
)

func (m cellMap) AddVec(v util.Vertex) {
	m[v] = true
}

func setup() {

}

func update() {

}

func draw() {

}
