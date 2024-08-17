package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		return
	}

	alive := 0
	universe := make([][]byte, n)
	for i := range universe {
		universe[i] = make([]byte, n)
	}

	//init
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if rand.Intn(2) == 1 {
				universe[i][j] = 'O'
				alive++
			} else {
				universe[i][j] = ' '
			}
		}
	}

	//last generation
	const generations = 50
	for i := 0; i < generations; i++ {
		universe = nextGeneration(universe, &alive)
		time.Sleep(150 * time.Millisecond)
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Generation #%d\n", i+1)
		fmt.Printf("Alive: %d\n", alive)
		printUniverse(universe)
	}

}

func nextGeneration(universe [][]byte, alive *int) [][]byte {
	n := len(universe)
	nextUniverse := make([][]byte, n)
	for i := 0; i < n; i++ {
		nextUniverse[i] = make([]byte, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			count := countAlive(universe, i, j)
			if universe[i][j] == 'O' {
				if count < 2 || count > 3 {
					nextUniverse[i][j] = ' '
					*alive--
				} else {
					nextUniverse[i][j] = 'O'
				}
			} else {
				if count == 3 {
					nextUniverse[i][j] = 'O'
					*alive++
				} else {
					nextUniverse[i][j] = ' '
				}
			}
		}
	}

	return nextUniverse
}

func countAlive(universe [][]byte, i, j int) int {
	n := len(universe)
	count := 0
	if universe[periodic(i-1, n)][periodic(j-1, n)] == 'O' {
		count++
	}
	if universe[periodic(i-1, n)][j] == 'O' {
		count++
	}
	if universe[periodic(i-1, n)][periodic(j+1, n)] == 'O' {
		count++
	}
	if universe[periodic(i+1, n)][periodic(j-1, n)] == 'O' {
		count++
	}
	if universe[periodic(i+1, n)][j] == 'O' {
		count++
	}
	if universe[periodic(i+1, n)][periodic(j+1, n)] == 'O' {
		count++
	}
	if universe[i][periodic(j-1, n)] == 'O' {
		count++
	}
	if universe[i][periodic(j+1, n)] == 'O' {
		count++
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
