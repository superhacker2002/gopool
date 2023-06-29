package options

import (
	"errors"
	"flag"
	"os"
)

var (
	ErrMultipleOptions = errors.New("multiple options are not allowed")
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
	foundFlag, err := validate([]bool{*lines, *characters, *words})

	if !foundFlag {
		*words = true
		files = os.Args[1:]
	}

	if err != nil {
		return Options{}, err
	}

	if len(files) == 0 {
		return Options{}, ErrNoFiles
	}

	return Options{
		L:     *lines,
		M:     *characters,
		W:     *words,
		Files: files,
	}, nil
}

func validate(flags []bool) (bool, error) {
	setFlags := 0
	for _, f := range flags {
		if f {
			setFlags++
		}
	}
	if setFlags > 1 {
		return true, ErrMultipleOptions
	}

	if setFlags == 0 {
		return false, nil
	}

	return true, nil
}
