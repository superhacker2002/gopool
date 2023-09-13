package main

import (
	"ex00/elasticclient"
	"ex00/parser"
	"flag"
	"log"
	"os"
)

func main() {
	dataPath := flag.String("data", "../../materials/data.csv", "path to csv file")
	indexName := flag.String("idxname", "places", "index name")
	docName := flag.String("docname", "place", "document name")
	flag.Parse()

	esClient, err := elasticclient.New()
	if err != nil {
		log.Fatalf("failed to create elastic search client: %v", err)
	}

	err = esClient.CreateIndex(*indexName)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = esClient.PutMapping(*indexName, *docName)
	if err != nil {
		log.Fatalf("%v", err)
	}

	file, err := os.Open(*dataPath)
	if err != nil {
		log.Fatalf("failed to open .csv file: %v", err)
	}

	records, err := parser.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to parse .csv file: %v", err)
	}

	err = esClient.SendData(records, *indexName)
	if err != nil {
		log.Fatalf("failed to load data to elastic search index: %v", err)
	}

}
