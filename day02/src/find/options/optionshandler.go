package options

import (
	"errors"
	"flag"
	"os"
)

var ErrWrongFlagComb = errors.New("flag -ext can not be used without -f")

type Options struct {
	Sl       bool
	D        bool
	F        bool
	Ext      string
	FilePath string
}

func New() (Options, error) {
	dir := flag.Bool("d", false, "./myFind -d /path/to/dir")
	symlink := flag.Bool("sl", false, "./myFind -Sl /path/to/dir")
	file := flag.Bool("f", false, "./myFind -f /path/to/dir")
	extension := flag.String("ext", "", "./myFind -f -ext 'go' /path/to/dir")

	flag.Parse()
	if !*file && *extension != "" {
		return Options{}, ErrWrongFlagComb
	}

	if !*dir && !*symlink && !*file {
		return Options{
			Sl:       true,
			D:        true,
			F:        true,
			Ext:      "",
			FilePath: os.Args[len(os.Args)-1],
		}, nil
	}

	return Options{
		Sl:       *symlink,
		D:        *dir,
		F:        *file,
		Ext:      *extension,
		FilePath: os.Args[len(os.Args)-1],
	}, nil
}
