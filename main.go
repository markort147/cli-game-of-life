package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var generations int
var rows int
var columns int
var pace int

func main() {
	flag.IntVar(&generations, "gens", 10, "Number of generations")
	flag.IntVar(&rows, "rows", 10, "Number of rows")
	flag.IntVar(&columns, "cols", 10, "Number of columns")
	flag.IntVar(&pace, "pace", 100, "Millis to wait for")
	flag.Parse()

	alive := 0
	universe := initUniverse(&alive)

	//last generation
	for i := 0; i < generations; i++ {
		universe = nextGeneration(universe, &alive)
		time.Sleep(time.Duration(pace) * time.Millisecond)
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Generation #%d\n", i+1)
		fmt.Printf("Alive: %d\n", alive)
		printUniverse(universe)
	}

}

func initUniverse(alive *int) [][]byte {
	universe := make([][]byte, rows)
	for i := range universe {
		universe[i] = make([]byte, columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if rand.Intn(2) == 1 {
				universe[i][j] = 'O'
				*alive++
			} else {
				universe[i][j] = ' '
			}
		}
	}
	return universe
}

func nextGeneration(universe [][]byte, alive *int) [][]byte {
	nextUniverse := make([][]byte, rows)
	var wg sync.WaitGroup

	for i := 0; i < rows; i++ {
		nextUniverse[i] = make([]byte, columns)
		for j := 0; j < columns; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()

				count := countAlive(universe, i, j)
				if universe[i][j] == 'O' {
					if count < 2 || count > 3 {
						nextUniverse[i][j] = ' '
					} else {
						nextUniverse[i][j] = 'O'
					}
				} else {
					if count == 3 {
						nextUniverse[i][j] = 'O'
					} else {
						nextUniverse[i][j] = ' '
					}
				}
			}(i, j)
		}
	}

	wg.Wait()

	*alive = 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if nextUniverse[i][j] == 'O' {
				*alive++
			}
		}
	}

	return nextUniverse
}

func countAlive(universe [][]byte, i, j int) int {
	count := 0
	neighbors := []struct {
		dx, dy int
	}{
		{-1, -1}, {-1, 0}, {-1, +1},
		{0, -1}, {0, +1},
		{+1, -1}, {+1, 0}, {+1, +1},
	}

	for _, neighbor := range neighbors {
		ni := periodic(i+neighbor.dx, rows)
		nj := periodic(j+neighbor.dy, columns)
		if universe[ni][nj] == 'O' {
			count++
		}
	}

	return count
}

func periodic(index int, n int) int {
	return (n + index) % n
}

func printUniverse(universe [][]byte) {
	for _, row := range universe {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}
