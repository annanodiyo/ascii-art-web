package ascii

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-web/printArt"
)

func Generate(text, banner string) (string, error) {
	// text = strings.ReplaceAll(text, "\\n", "\n")

	if banner == "" {
		banner = "standard"
	}

	filepath := "banners/" + banner + ".txt"

	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("banner not found")
	}
	text = strings.ReplaceAll(text, "\\n", "\n")

	lines := strings.Split(string(data), "\n")

	// text = strings.ReplaceAll(text, "\\n", "\n")

	words := strings.Split(text, "\n")

	var result strings.Builder

	for i, word := range words {
		art := printArt.GenerateArt(word, lines)
		result.WriteString(art)

		if i != len(words)-1 {
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}
