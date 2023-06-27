package main

import (
	"day01/internal/comparators/dbcomparator"
	"day01/internal/dbreaders/jsonreader"
	"day01/internal/dbreaders/xmlreader"
	"day01/internal/entity"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

const usage string = "./compareDB --old original_database.xml --new stolen_database.json"

type DBReader interface {
	Read(reader io.Reader) (entity.CakeRecipes, error)
}

func main() {
	oldFileName := flag.String("old", "", usage)
	newFileName := flag.String("new", "", usage)
	flag.Parse()

	if *oldFileName == "" || *newFileName == "" {
		log.Fatal("one of the files was not provided")
	}
	oldFileReader := chooseHandlers(*oldFileName)
	newFileReader := chooseHandlers(*newFileName)

	if oldFileReader == nil || newFileReader == nil {
		log.Fatal("file must have .json or .xml extension")
	}

	file, err := os.Open(*oldFileName)

	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	oldRecipes, err := oldFileReader.Read(file)
	if err != nil {
		log.Fatalf("failed to read the file: %v", err)
	}

	file, err = os.Open(*newFileName)

	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	newRecipes, err := newFileReader.Read(file)
	if err != nil {
		log.Fatalf("failed to read the file: %v", err)
	}

	fmt.Println(dbcomparator.Compare(oldRecipes, newRecipes))
}

func chooseHandlers(fileName string) DBReader {
	if path.Ext(fileName) == ".json" {
		return jsonreader.JsonReader{}
	} else if path.Ext(fileName) == ".xml" {
		return xmlreader.XmlReader{}
	}
	return nil
}
