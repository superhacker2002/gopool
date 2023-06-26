package main

import (
	"day01/internal/converters/jsonconverter"
	"day01/internal/converters/xmlconverter"
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

const (
	JSONFile = 1
	XMLFile  = 2
)

type DBReader interface {
	Read(reader io.Reader) (entity.CakeRecipes, error)
}

type EntityConverter interface {
	Convert(recipes entity.CakeRecipes) (string, error)
}

func main() {
	fileName := flag.String("f", "", "./readDB -f original_database.xml")
	flag.Parse()

	if *fileName == "" {
		log.Fatal("file name was not provided")
	}

	var (
		reader    DBReader
		converter EntityConverter
	)
	reader, converter = chooseHandlers(*fileName)
	if reader == nil || converter == nil {
		log.Fatal("file must have .json or .xml extension")
	}

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}

	recipes, err := reader.Read(file)
	if err != nil {
		log.Fatalf("failed to read the file: %v", err)
	}

	fmt.Println(converter.Convert(recipes))
}

func chooseHandlers(fileName string) (DBReader, EntityConverter) {
	if path.Ext(fileName) == ".json" {
		return jsonreader.JsonReader{}, xmlconverter.XmlConverter{}
	} else if path.Ext(fileName) == ".xml" {
		return xmlreader.XmlReader{}, jsonconverter.JsonConverter{}
	}
	return nil, nil
}
