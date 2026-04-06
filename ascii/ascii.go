package ascii

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func GenerateAscii(text string, banner string) (string, error) {
	filePath := "banners/" + banner + ".txt"

	file, err := os.Open(filePath)
	if err != nil {
		return "", errors.New("banner not found")
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	inputLines := strings.Split(text, "\n")

	var result strings.Builder

	for _, line := range inputLines {
		if line == "" {
			result.WriteString("\n")
			continue
		}

		for i := 0; i < 8; i++ {
			for _, char := range line {
				if char < 32 || char > 126 {
					return "", errors.New("invalid character")
				}

				index := (int(char) - 32) * 9
				result.WriteString(lines[index+i])
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}