package main

import (
	"flag"
	"fmt"
	"github.com/inancgumus/screen"
	"main/config"
	"main/universe"
	"strconv"
	"time"
)

func main() {
	defineUsage()
	cfg := config.ParseFlags()
	timeline := universe.NewTimeline(cfg.Generations, cfg.Rows, cfg.Columns, cfg.Density)
	runApp(timeline, cfg.Generations, cfg.Pace)
}

func defineUsage() {
	flag.CommandLine.Usage = func() {
		fmt.Print(`The Game of Life is a cellular automaton devised by British mathematician John Horton Conway in 1970.
This application starts with a random configuration and evolves according to these rules:
    
    1. Any live cell with fewer than two live neighbors dies (underpopulation).
    2. Any live cell with two or three live neighbors survives to the next generation.
    3. Any live cell with more than three live neighbors dies (overpopulation).
    4. Any dead cell with exactly three live neighbors becomes a live cell (reproduction).

For more information, visit: https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life

You can launch the app with the following optional arguments:
`)
		flag.PrintDefaults()
		fmt.Print(`    
After launching, the app will compute and print the starting configuration of all generations.
You can navigate through generations with these commands:    
    - 'g' followed by an optional integer (default is 0, no spaces)
        Display the specified generation.
    - 'n' followed by an optional integer (default is 1, no spaces)
        Navigate forward by the specified number of generations.
    - 'p' followed by an optional integer (default is 1, no spaces)
        Navigate backward by the specified number of generations.
    - 'a'
        Activate auto mode, where generations are displayed sequentially
        at the pace set by the user.
    - 'r'
        Display the initial configuration.
    - 'q'
        Quit the application.
    
Enjoy! 😀`)
	}
}

func runApp(timeline []universe.Universe, generations int, pace int) {
	currGen := 0
	autoMode := false

	for {
		clearAndRender(&timeline[currGen])

		if autoMode {
			if currGen < generations {
				currGen++
				time.Sleep(time.Duration(pace) * time.Millisecond)
			} else {
				//Disable auto mode at the end
				autoMode = false
			}
		} else {
			fmt.Print("Enter command [(n)ext (p)revious (g)eneration (q)uit (a)uto (r)eset]: ")
			var input string
			_, err := fmt.Scanln(&input)
			if err != nil || len(input) == 0 {
				continue
			}

			switch input[0] {
			case 'n':
				arg := extractArg(input, 1)
				currGen = min(currGen+arg, generations)
			case 'p':
				arg := extractArg(input, 1)
				currGen = max(currGen-arg, 0)
			case 'g':
				arg := extractArg(input, 0)
				currGen = min(arg, generations)
			case 'q':
				return
			case 'a':
				autoMode = true
			case 'r':
				currGen = 0
			default:
				continue
			}
		}
	}
}

// clearAndRender clears the screen and renders the current universe state.
func clearAndRender(currUniverse *universe.Universe) {
	screen.Clear()
	screen.MoveTopLeft()
	currUniverse.Print()
}

// extractArg extracts an argument from the input or returns a default value.
func extractArg(input string, defaultValue int) int {
	if len(input) > 1 {
		arg, err := strconv.Atoi(input[1:])
		if err != nil {
			return defaultValue
		}
		return arg
	}
	return defaultValue
}
