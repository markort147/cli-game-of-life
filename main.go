package main

import (
	"flag"
	"fmt"
	"github.com/inancgumus/screen"
	"log"
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

	timeline := make([]universe.Universe, *generations+1)
	timeline[0].Init(*rows, *columns, *density)
	for gen := 1; gen < *generations+1; gen++ {
		timeline[gen] = timeline[gen-1].NextGeneration()
	}

	screen.Clear()
	exit := false
	currGen := 0
	autoMode := false
	for !exit {
		screen.MoveTopLeft()
		timeline[currGen].Print()

		if autoMode {
			if currGen < *generations {
				currGen++
				time.Sleep(time.Duration(*pace) * time.Millisecond)
			} else {
				autoMode = false
			}
		} else {
			fmt.Print("Enter command [(n)ext (p)revious (q)uit (a)uto (r)eset]: ")
			var command string
			_, err := fmt.Scanf("%s\n", &command)
			if err != nil {
				log.Fatal(err)
			}

			switch command {
			case "q":
				exit = true
			case "n":
				if currGen < *generations {
					currGen++
				}
			case "p":
				if currGen > 0 {
					currGen--
				}
			case "a":
				autoMode = true
			case "r":
				currGen = 0
			}
		}
	}
}
