package multiplex

import (
	"testing"
)

func TestMultiplexAllValuesReceived(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})
	result := Multiplex(ch1, ch2, ch3)

	values := []interface{}{"value1", 42, true}
	expectedCount := len(values) * 3

	go func() {
		for _, value := range values {
			ch1 <- value
			ch2 <- value
			ch3 <- value
		}
		close(ch1)
		close(ch2)
		close(ch3)
	}()

	receivedCount := 0
	for value := range result {
		receivedCount++
		if !contains(values, value) {
			t.Errorf("Unexpected value received: %v", value)
		}
	}

	if receivedCount != expectedCount {
		t.Errorf("Expected %d values, received %d", expectedCount, receivedCount)
	}
}

func contains(slice []interface{}, value interface{}) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
