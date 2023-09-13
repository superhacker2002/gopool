package crawler

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestReadBody_Success(t *testing.T) {
	expected := "Hello, World!"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expected)
	}))
	defer server.Close()

	body := readBody(server.URL)
	if body != expected {
		t.Errorf("Expected body %q, but got %q", expected, body)
	}
}

func TestCrawlWeb_Success(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan string)
	res := CrawlWeb(ctx, ch)

	url := "http://example.com"
	ch <- url

	time.Sleep(100 * time.Millisecond)

	cancel()

	select {
	case body := <-res:
		if body != readBody(url) {
			t.Errorf("Expected body %q, but got %q", readBody(url), body)
		}
	case <-time.After(1 * time.Second):
		t.Error("Expected res channel to receive a result, but it didn't")
	}
}
