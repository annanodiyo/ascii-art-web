package printArt

import (
	"strings"
)

const height = 8

func GenerateArt(text string, lines []string) string {
	var result strings.Builder
	for _, line := range strings.Split(text, "\n") {
		output := make([]string, height)

		for _, c := range line {
			var block []string
			ascii := int(c)

			if ascii < 32 || ascii > 126 {
				ascii = 32
			}

			index := (ascii - 32) * 9
			if index+height <= len(lines) {
				block = lines[index : index+height]
			} else {
				block = lines[0:height]
			}

			for j := 0; j < height; j++ {
				output[j] += block[j]
			}
		}

		for _, out := range output {
			// fmt.Println(out)
			result.WriteString(out + "\n")
		}

		if line == "" {
			// fmt.Println()
			result.WriteString("\n")
		}
	}

	return result.String()
}
