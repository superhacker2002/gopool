package main

import (
	"log"
	"myFind/file"
	"myFind/options"
)

func main() {
	opts, err := options.New()
	if err != nil {
		log.Fatalf("failed to get command line arguments: %v", err)
	}

	err = file.Find(opts)
	if err != nil {
		log.Fatalf("failed to walk through the directory: %v", err)
	}

}
