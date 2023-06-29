package count

import (
	"bufio"
	"io"
)

func Lines(reader io.Reader) (int, error) {
	fileScanner := bufio.NewScanner(reader)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}

	if err := fileScanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}
