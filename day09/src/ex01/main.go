package main

import (
	"context"
	"ex01/crawler"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	killSig := make(chan os.Signal)
	signal.Notify(killSig, syscall.SIGTERM, os.Interrupt)
	go func() {
		<-killSig
		fmt.Println("\nprogram stopped")
		cancel()
	}()

	in := make(chan string)

	go func() {
		select {
		case <-ctx.Done():
			return
		case in <- "http://example.com/":
		}
	}()

	data := crawler.CrawlWeb(ctx, in)

	for {
		select {
		case <-ctx.Done():
			return
		case str := <-data:
			fmt.Println(str)
		}
	}
}
