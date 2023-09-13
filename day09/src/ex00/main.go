package main

import (
	"ex00/sorter"
	"fmt"
)

func main() {
	ch := sorter.SleepSort([]int{1, 4, 7, 2, 35, 9})

	for num := range ch {
		fmt.Println(num)
	}
}
