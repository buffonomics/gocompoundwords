package util

import (
	"bufio"
	"os"
)

func LoadStrings(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		if l != "" {
			lines = append(lines, l)
		}
	}
	return lines, scanner.Err()
}
