package main

import (
	"ex00/arithmetic"
	"fmt"
	"log"
)

func main() {
	el, err := arithmetic.GetElement([]int{1, 2, 3, 4, 5}, 1)
	if err != nil {
		log.Fatal("error while getting element by index:", err)
	}
	fmt.Println(el)
}
