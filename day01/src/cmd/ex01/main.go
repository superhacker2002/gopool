package main

import (
	"day01/internal/entity"
	"flag"
	"io"
	"log"
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

}
