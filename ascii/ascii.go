package ascii

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-web/printArt"
)

// func Generate(text, banner string) (string, error) {
// 	if banner == "" {
// 		return "", fmt.Errorf("banner content is empty")
// 	}
// 	lines := strings.Split(banner, "\n")
// 	return printArt.GenerateArt(text, lines), nil
// }

func Generate(text, banner string) (string, error) {
	filepath := "../banners/" + banner + ".txt"
	data, err := os.ReadFile(filepath)
	// fmt.Println(data)
	if err != nil {
		return "", fmt.Errorf("banner not found")
	}
	lines := strings.Split(string(data), "\n")
	// fmt.Println(data)
	return printArt.GenerateArt(text, lines), nil
}
