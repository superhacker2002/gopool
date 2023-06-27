package main

import (
	"day01/internal/comparators/fbcomparator"
	"flag"
	"fmt"
	"log"
	"os"
)

const usage = "./compareFS --old snapshot1.txt --new snapshot2.txt"

func main() {
	oldFileName := flag.String("old", "", usage)
	newFileName := flag.String("new", "", usage)
	flag.Parse()

	if *oldFileName == "" || *newFileName == "" {
		log.Fatal("one of the files was not provided")
	}

	oldFile, err := os.Open(*oldFileName)

	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer oldFile.Close()

	newFile, err := os.Open(*newFileName)

	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer newFile.Close()

	result, err := fbcomparator.Compare(oldFile, newFile)
	if err != nil {
		log.Fatalf("failed to compare files: %v", err)
	}

	fmt.Print(result)
}
