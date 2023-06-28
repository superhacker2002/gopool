package options

import "flag"

type Options struct {
	L bool
	M bool
	W bool
}

func New() (Options, error) {
	lines := flag.Bool("l", false, "./myWc -l input.txt")
	characters := flag.Bool("m", false, "./myWc -m input.txt")
	words := flag.Bool("w", false, "./myWc -w input.txt")

}
