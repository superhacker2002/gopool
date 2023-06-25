package main

import (
	"day01/converter"
	"day01/dbreaders/jsonreader"
	"day01/dbreaders/xmlreader"
	"day01/entity"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

type DBReader interface {
	Read(reader io.Reader) (entity.CakeRecipes, error)
}

func main() {
	fileName := flag.String("f", "", "./readDB -f original_database.xml")
	flag.Parse()

	if *fileName == "" {
		log.Fatal("file name was not provided")
	}

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}

	reader := chooseReader(*fileName)

	recipes, err := reader.Read(file)
	if err != nil {
		log.Fatalf("error while reading file: %v", err)
	}
	xmlRecipe, err := converter.ToXml(recipes)
	if err != nil {
		log.Fatalf("failed to connvert structure to JSON: %v", err)
	}

	fmt.Println(xmlRecipe)
}

func chooseReader(fileName string) DBReader {
	if path.Ext(fileName) == ".json" {
		return jsonreader.JsonReader{}
	}
	if path.Ext(fileName) == ".xml" {
		return xmlreader.XmlReader{}
	}
	return nil
}
