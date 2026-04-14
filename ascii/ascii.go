package ascii

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-web/printArt"
)

func Generate(text, banner string) (string, error) {

	// Default banner
	if banner == "" {
		banner = "standard"
	}

	// FIX: Correct path (REMOVED "../")
	filepath := "banners/" + banner + ".txt"

	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("banner not found")
	}

	lines := strings.Split(string(data), "\n")

	return printArt.GenerateArt(text, lines), nil
}