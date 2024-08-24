package main

import (
	"flag"
	"fmt"
	"main/universe"
	"time"
)

func main() {
	generations := flag.Int("gens", 20, "Number of generations")
	rows := flag.Int("rows", 15, "Number of rows")
	columns := flag.Int("cols", 50, "Number of columns")
	pace := flag.Int("pace", 50, "Millis to wait for")
	density := flag.Float64("dens", 0.5, "Initial density of aliveCounter cells")
	flag.Parse()

	var uni universe.Universe
	uni.Init(*rows, *columns, *density)

	for range *generations + 1 {
		fmt.Print("\033[H\033[2J")
		uni.Print()
		uni = uni.NextGeneration()
		time.Sleep(time.Duration(*pace) * time.Millisecond)
	}
}
