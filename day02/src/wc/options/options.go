package options

import (
	"errors"
	"flag"
	"os"
)

var (
	ErrMultipleOptions = errors.New("multiple options are not allowed")
	ErrNoOptions       = errors.New("at least one option should be specified")
	ErrNoFiles         = errors.New("files were not provided")
)

type Options struct {
	L     bool
	M     bool
	W     bool
	Files []string
}

func New() (Options, error) {
	lines := flag.Bool("l", false, "./myWc -l input.txt")
	characters := flag.Bool("m", false, "./myWc -m input.txt")
	words := flag.Bool("w", false, "./myWc -w input.txt")

	flag.Parse()

	files := os.Args[2:]
	err := validate([]bool{*lines, *characters, *words}, files)
	if err != nil {
		return Options{}, err
	}

	return Options{
		L:     *lines,
		M:     *characters,
		W:     *words,
		Files: files,
	}, nil
}

func validate(flags []bool, files []string) error {
	setFlags := 0
	for _, f := range flags {
		if f {
			setFlags++
		}
	}
	if setFlags > 1 {
		return ErrMultipleOptions
	}

	if setFlags == 0 {
		flags[2] = true
	}

	if len(files) == 0 {
		return ErrNoFiles
	}

	return nil
}
