package main

import (
	"flag"
	"github.com/inancgumus/screen"
	"main/universe"
	"time"
)

func main() {
	generations := flag.Int("gens", 20, "Number of generations")
	rows := flag.Int("rows", 15, "Number of rows")
	columns := flag.Int("cols", 50, "Number of columns")
	pace := flag.Int("pace", 50, "Time pace between generations (millis)")
	density := flag.Float64("dens", 0.5, "Initial density of random cells")
	flag.Parse()

	var uni universe.Universe
	uni.Init(*rows, *columns, *density)

	screen.Clear()
	for range *generations + 1 {
		screen.MoveTopLeft()
		uni.Print()
		uni = uni.NextGeneration()
		time.Sleep(time.Duration(*pace) * time.Millisecond)
	}
}
