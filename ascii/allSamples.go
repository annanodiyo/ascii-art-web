package ascii

import "strings"

func AllSamples(text string) (string, error) {
	banners := []string{"standard", "shadow", "thinkertoy"}

	var result strings.Builder
	for _, banner := range banners {
		art, err := Generate(text, banner)
		if err != nil {
			return "", err
		}
		result.WriteString("===" + banner + "===\n")
		result.WriteString(art)
		result.WriteString("\n\n")
	}
	return result.String(), nil
}
