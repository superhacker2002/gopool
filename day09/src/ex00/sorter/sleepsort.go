package sorter

import (
	"sync"
	"time"
)

func SleepSort(nums []int) chan int {
	ch := make(chan int, len(nums))
	defer close(ch)

	var wg sync.WaitGroup
	wg.Add(len(nums))
	for _, num := range nums {
		num := num
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(num) * time.Millisecond)
			ch <- num
		}()
	}
	wg.Wait()
	return ch
}
