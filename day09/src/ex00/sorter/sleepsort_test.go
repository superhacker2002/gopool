package sorter

import "testing"

func TestSleepSort(t *testing.T) {
	nums := []int{10, 5, 8, 3, 2, 7, 1, 6, 9, 4}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch := SleepSort(nums)

	sorted := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		sorted[i] = <-ch
	}

	for i := 0; i < len(nums); i++ {
		if sorted[i] != expected[i] {
			t.Errorf("Expected %d at index %d, but got %d", expected[i], i, sorted[i])
		}
	}
}
