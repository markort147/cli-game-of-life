package config

import "flag"

type Config struct {
	Generations int
	Rows        int
	Columns     int
	Pace        int
	Density     float64
}

// ParseFlags parses the command-line flags and returns a Config.
func ParseFlags() Config {
	generations := flag.Int("gens", 50, "Number of generations")
	rows := flag.Int("rows", 20, "Number of rows")
	columns := flag.Int("cols", 40, "Number of columns")
	pace := flag.Int("pace", 150, "Time pace between generations in auto mode (milliseconds)")
	density := flag.Float64("dens", 0.4, "Initial density of random cells")
	flag.Parse()

	return Config{
		Generations: *generations,
		Rows:        *rows,
		Columns:     *columns,
		Pace:        *pace,
		Density:     *density,
	}
}
