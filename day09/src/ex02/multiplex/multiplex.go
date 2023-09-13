package multiplex

import "sync"

func Multiplex(channels ...<-chan interface{}) chan interface{} {
	out := make(chan any)

	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, c := range channels {
		c := c
		go func() {
			for n := range c {
				out <- n
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
