package count

import (
	"bufio"
	"io"
	"unicode/utf8"
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

func Characters(reader io.Reader) (int, error) {
	fileScanner := bufio.NewScanner(reader)
	charsCount := 0

	for fileScanner.Scan() {
		charsCount += utf8.RuneCountInString(fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		return 0, err
	}

	return charsCount, nil
}

func Words(reader io.Reader) (int, error) {
	fileScanner := bufio.NewScanner(reader)
	wordsCount := 0

	fileScanner.Split(bufio.ScanWords)
	for fileScanner.Scan() {
		wordsCount++
	}

	if err := fileScanner.Err(); err != nil {
		return 0, err
	}

	return wordsCount, nil
}
