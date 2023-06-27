package fbcomparator

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"io"
)

func Compare(first io.Reader, second io.Reader) (string, error) {
	outputStr := ""
	firstFiles, err := readFile(first)
	if err != nil {
		return "", err
	}

	return outputStr + compareLines(firstFiles, second), nil
}

func readFile(reader io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func compareLines(first []string, second io.Reader) string {
	var inSecond []string
	outputStr := ""
	scanner := bufio.NewScanner(second)
	for scanner.Scan() {
		if !slices.Contains(first, scanner.Text()) {
			outputStr += fmt.Sprintf("ADDED %s\n", scanner.Text())
		} else {
			inSecond = append(inSecond, scanner.Text())
		}
	}

	for _, file := range first {
		if !slices.Contains(inSecond, file) {
			outputStr += fmt.Sprintf("REMOVED %s\n", file)
		}
	}
	return outputStr
}
