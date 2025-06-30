package ascii

import "strings"

func RenderAscii(input string, banner map[rune][]string) string {
	var result strings.Builder
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}

		var outputRows [8]string
		for _, ch := range line {
			if val, ok := banner[ch]; ok {
				for i := 0; i < 8; i++ {
					outputRows[i] += val[i]
				}
			}
		}

		for i := 0; i < 8; i++ {
			result.WriteString(outputRows[i] + "\n")
		}
	}

	return result.String()
}
