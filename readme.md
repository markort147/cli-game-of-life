# Conway's Game of Life

Conway's Game of Life is a cellular automaton devised by British mathematician John Horton Conway in 1970. This project
is a terminal-based implementation of the Game of Life in Go.

## Overview

The Game of Life is a zero-player game, meaning its evolution is determined by its initial state, requiring no further
input. The universe of the Game of Life is a grid of cells, where each cell can either be alive or dead. Cells evolve
over generations based on the number of alive neighbors according to the following rules:

1. Any live cell with fewer than two live neighbors dies (underpopulation).
2. Any live cell with two or three live neighbors survives to the next generation.
3. Any live cell with more than three live neighbors dies (overpopulation).
4. Any dead cell with exactly three live neighbors becomes a live cell (reproduction).
For more information on the Game of Life, you can visit its [Wikipedia page](https://en.wikipedia.org/wiki/Conway's_Game_of_Life).

## Features

- Generates a random initial configuration of live and dead cells based on user-defined density.
- Simulates the Game of Life over multiple generations.
- Allows navigation through generations, both forward and backward.
- Supports auto mode for continuous simulation at a user-defined pace.

## Installation

### 1. Clone the repository:

``` bash
git clone https://github.com/yourusername/game-of-life.git
cd game-of-life
```

### 2. Install dependencies:

Make sure you have [Go installed](https://golang.org/doc/install) on your system. This project uses Go modules, so
dependencies will be handled automatically.

### 3. Build the project:

``` bash
go build -o game-of-life
```

### 4. Run the project:

``` bash
./game-of-life [options]
```

## Usage

### Command-Line Options

You can configure the simulation using command-line flags. Here are the available options:

- `-gens`: The number of generations to simulate. (Default: 50)
- `-rows`: The number of rows in the universe. (Default: 20)
- `-cols`: The number of columns in the universe. (Default: 40)
- `-pace`: The time delay (in milliseconds) between generations in auto mode. (Default: 150ms)
- `-dens`: The initial density of live cells as a float between 0 and 1. (Default: 0.4)

Example:

``` bash
./game-of-life -gens 100 -rows 30 -cols 50 -pace 200 -dens 0.5
```

### In-Game Commands

Once the simulation has started, you can control the simulation using these commands:

- n: Move to the next generation. Optionally, follow the n with a number (e.g., n5) to skip ahead by that many
  generations.
- p: Move to the previous generation. Optionally, follow the p with a number (e.g., p3) to go back by that many
  generations.
- g: Jump to a specific generation. Follow the g with the generation number (e.g., g10).
- r: Reset the simulation to the initial configuration (generation 0).
- a: Activate auto mode to automatically move through generations at the specified pace.
- q: Quit the simulation.

### Example Usage

``` bash
./game-of-life -gens 100 -rows 30 -cols 50 -pace 200 -dens 0.5
```

Once running, use the following commands:

- Enter n5 to skip to the 5th generation.
- Enter p2 to go back 2 generations.
- Enter a to activate auto mode.
- Enter q to quit.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your improvements or new features.

### Steps to Contribute

1. Fork the repository.
2. Create a new branch.
3. Make your changes.
4. Submit a pull request.

## Contact

If you have any questions or feedback, feel free to open an issue or contact me at [marco93romano@gmail.com](mailto://marco93romano@gmail.com).

