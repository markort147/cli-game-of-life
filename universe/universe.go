package universe

import (
	"fmt"
	"math/rand"
	"sync"
)

type Universe struct {
	cells        [][]byte
	aliveCounter int
	generation   int
}

func (u *Universe) Init(rows, columns int, density float64) {
	u.cells = make([][]byte, rows)
	for row := range rows {
		u.cells[row] = make([]byte, columns)
		for col := range columns {
			var cellValue byte
			if rand.Float64() < density {
				cellValue = 'O'
				u.aliveCounter++
			} else {
				cellValue = ' '
			}
			u.cells[row][col] = cellValue
		}
	}
}

func (u *Universe) NextGeneration() Universe {
	var nextUniverse Universe
	nextUniverse.generation = u.generation + 1
	nextUniverse.cells = make([][]byte, len(u.cells))

	var wg sync.WaitGroup
	for row := range u.cells {
		nextUniverse.cells[row] = make([]byte, len(u.cells[row]))
		for col := range u.cells[row] {
			wg.Add(1)

			go func(row, col int) {
				defer wg.Done()

				aliveNeighbors := u.AliveNeighbors(row, col)
				var cellValue byte
				if u.cells[row][col] == 'O' {
					if aliveNeighbors < 2 || aliveNeighbors > 3 {
						cellValue = ' '
					} else {
						cellValue = 'O'
					}
				} else {
					if aliveNeighbors == 3 {
						cellValue = 'O'
					} else {
						cellValue = ' '
					}
				}
				nextUniverse.cells[row][col] = cellValue
			}(row, col)
		}
	}
	wg.Wait()

	for row := range u.cells {
		for col := range u.cells[row] {
			if nextUniverse.cells[row][col] == 'O' {
				nextUniverse.aliveCounter++
			}
		}
	}

	return nextUniverse
}

func (u *Universe) AliveNeighbors(i, j int) int {
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
		if u.cells[ni][nj] == 'O' {
			count++
		}
	}

	return count
}

func (u *Universe) Print() {
	fmt.Printf("Generation #%d\n", u.generation)
	fmt.Printf("Alive: %d\n", u.aliveCounter)
	for _, row := range u.cells {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}
