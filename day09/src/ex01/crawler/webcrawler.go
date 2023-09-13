package crawler

import (
	"context"
	"io"
	"log"
	"net/http"
)

const maxGoroutines = 8

func CrawlWeb(ctx context.Context, ch chan string) chan string {
	res := make(chan string)
	sem := make(chan struct{}, maxGoroutines)

	go func() {
		select {
		case <-ctx.Done():
			return
		case url := <-ch:
			sem <- struct{}{}
			go func() {
				select {
				case <-ctx.Done():
					return
				default:
					res <- readBody(url)
				}
				<-sem
			}()
		}
	}()

	return res

}

func readBody(url string) string {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("server responded with status code %s", resp.Status)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(bodyBytes)
}
