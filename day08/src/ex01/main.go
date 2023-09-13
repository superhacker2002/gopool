package main

import "ex01/description"

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

func main() {
	description.DescribePlant(AnotherUnknownPlant{
		FlowerColor: 10,
		LeafType:    "lanceolate",
		Height:      15,
	})

	description.DescribePlant(UnknownPlant{
		FlowerType: "rose",
		LeafType:   "good",
		Color:      255,
	})
}
