package main

import (
	"ex03/multiplex"
	"fmt"
)

func main() {
	ch1 := genChannel(100)
	ch2 := genChannel(2)
	ch3 := genChannel(3333)
	out := multiplex.Multiplex(ch1, ch2, ch3)

	for v := range out {
		fmt.Println(v)
	}
}

func genChannel(data any) <-chan any {
	ch := make(chan any)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- data
		}
		close(ch)
	}()

	return ch
}
