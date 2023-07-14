package gotools

import (
	"bufio"
	"os"
	"regexp"
)

// ej files/file.txt, `Input: input\.(\w+)\(\)`
func GetNamesFromFileAndPattern(file_path, pattern string) (out []string) {

	file, err := os.Open(file_path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inputPattern := regexp.MustCompile(pattern)

	for scanner.Scan() {
		line := scanner.Text()
		match := inputPattern.FindStringSubmatch(line)
		if match != nil {
			out = append(out, match[1])
		}
	}

	return
}
