package file

import (
	"fmt"
	"myFind/options"
	"os"
	"path/filepath"
)

func Find(opts options.Options) error {
	err := filepath.Walk(opts.FilePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if os.ModePerm != 0 {
			if info.IsDir() && opts.D {
				fmt.Println(path)
			}

			if info.Mode().Type() == os.ModeSymlink && opts.Sl {
				processSymLink(path, info)
			}

			if info.Mode().IsRegular() && opts.F {
				processFile(path, opts.Ext)
			}

		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func processFile(path string, ext string) {
	if ext != "" {
		if filepath.Ext(path) == "."+ext {
			fmt.Println(path)
		}
	} else {
		fmt.Println(path)
	}
}

func processSymLink(path string, info os.FileInfo) {
	location, err := os.Readlink(info.Name())
	if err != nil {
		location = "[broken]"
	}
	fmt.Println(path, "->", location)
}
