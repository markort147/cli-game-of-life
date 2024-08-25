package universe

import (
	"fmt"
	"math/rand"
)

const deadMarker = ' '
const aliveMarker = 'O'

type cell struct {
	marker byte
}

func (c *cell) isAlive() bool {
	return c.marker == aliveMarker
}

func (c *cell) arise() {
	c.marker = aliveMarker
}

func (c *cell) die() {
	c.marker = deadMarker
}

type Universe struct {
	cells        [][]cell
	aliveCounter int
	generation   int
}

func (u *Universe) Init(rows, columns int, density float64) {
	u.cells = make([][]cell, rows)
	for row := range rows {
		u.cells[row] = make([]cell, columns)
		for col := range columns {
			if rand.Float64() < density {
				u.cells[row][col].arise()
				u.aliveCounter++
			} else {
				u.cells[row][col].die()
			}
		}
	}
}

func (u *Universe) NextGeneration() Universe {
	var nextUniverse Universe
	nextUniverse.generation = u.generation + 1
	nextUniverse.cells = make([][]cell, len(u.cells))

	for row := range u.cells {
		nextUniverse.cells[row] = make([]cell, len(u.cells[row]))
		for col := range u.cells[row] {
			aliveNeighbors := u.aliveNeighbors(row, col)
			if u.cells[row][col].isAlive() {
				if aliveNeighbors < 2 || aliveNeighbors > 3 {
					nextUniverse.cells[row][col].die()
				} else {
					nextUniverse.cells[row][col].arise()
					nextUniverse.aliveCounter++
				}
			} else {
				if aliveNeighbors == 3 {
					nextUniverse.cells[row][col].arise()
					nextUniverse.aliveCounter++
				} else {
					nextUniverse.cells[row][col].die()
				}
			}

		}
	}

	return nextUniverse
}

func (u *Universe) aliveNeighbors(i, j int) int {
	rows := len(u.cells)
	cols := len(u.cells[0])

	neighbors := []struct {
		dx, dy int
	}{
		{-1, -1}, {-1, 0}, {-1, +1},
		{0, -1}, {0, +1},
		{+1, -1}, {+1, 0}, {+1, +1},
	}

	count := 0
	for _, neighbor := range neighbors {
		ni := (rows + i + neighbor.dx) % rows
		nj := (cols + j + neighbor.dy) % cols
		if u.cells[ni][nj].isAlive() {
			count++
		}
	}

	return count
}

func (u *Universe) Print() {
	fmt.Printf("Generation #%d\n", u.generation)
	fmt.Printf("Alive: %d\n", u.aliveCounter) //todo fix print of alive between generations
	for _, row := range u.cells {
		for _, cell := range row {
			fmt.Print(string(cell.marker))
		}
		fmt.Println()
	}
}
