package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"myWc/count"
	"myWc/options"
	"os"
)

type countFunc func(reader io.Reader) (int, error)

func main() {
	opts, err := options.New()
	if err != nil {
		log.Fatalf("failed to get command line arguments: %v", err)
	}

	start(opts)
}

func start(opts options.Options) {
	g := new(errgroup.Group)
	ch := make(chan string, len(opts.Files))
	if opts.L {
		processFiles(opts, count.Lines, g, ch)
	}
	if opts.M {
		processFiles(opts, count.Characters, g, ch)
	}
	if opts.W {
		processFiles(opts, count.Words, g, ch)
	}

	if err := g.Wait(); err != nil {
		log.Fatalf("failed to process files: %v", err)
	}
	close(ch)

	for output := range ch {
		fmt.Println(output)
	}

}

func processFiles(opts options.Options, count countFunc, g *errgroup.Group, ch chan string) {
	for _, fileName := range opts.Files {
		fileName := fileName
		g.Go(func() error {
			file, err := os.Open(fileName)
			if err != nil {
				return err
			}
			number, err := count(file)
			ch <- fmt.Sprintf("%d\t%s", number, fileName)
			return err
		})
	}
}
