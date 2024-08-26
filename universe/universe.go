package universe

import (
	"fmt"
	"math/rand"
)

// Constants representing the cell states
const (
	deadMarker  = ' '
	aliveMarker = '𖡹'
)

// Cell represent a single cell in the universe
type cell struct {
	marker rune
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

// Universe represents the grid of cells and related generation information
type Universe struct {
	cells        [][]cell
	aliveCounter int
	generation   int
}

// NewTimeline initializes a slice of universes, each representing a generation.
func NewTimeline(generations int, rows int, columns int, density float64) []Universe {
	timeline := make([]Universe, generations+1)
	timeline[0].init(rows, columns, density)
	for gen := 1; gen < generations+1; gen++ {
		timeline[gen] = timeline[gen-1].nextGeneration()
	}
	return timeline
}

// init initializes the universe with a random configuration based on the density parameter
func (u *Universe) init(rows, columns int, density float64) {
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

// nextGeneration computes the next state of the universe based on the current state.
func (u *Universe) nextGeneration() Universe {
	next := Universe{
		cells:      make([][]cell, len(u.cells)),
		generation: u.generation + 1,
	}

	for row := range u.cells {
		next.cells[row] = make([]cell, len(u.cells[row]))
		for col := range u.cells[row] {
			aliveNeighbors := u.aliveNeighbors(row, col)
			if u.cells[row][col].isAlive() {
				if aliveNeighbors < 2 || aliveNeighbors > 3 {
					next.cells[row][col].die()
				} else {
					next.cells[row][col].arise()
					next.aliveCounter++
				}
			} else {
				if aliveNeighbors == 3 {
					next.cells[row][col].arise()
					next.aliveCounter++
				} else {
					next.cells[row][col].die()
				}
			}

		}
	}

	return next
}

// aliveNeighbors counts the number of alive neighbors for a given cell
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

// Print outputs the current state of the universe to the console
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
