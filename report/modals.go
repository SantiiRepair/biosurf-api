package report

import (
	"strings"
)

func isProfanity(text string) bool {
	text = strings.ToLower(text)

	for _, word := range wordsList {
		if strings.Contains(text, word) {
			return true
		}
	}

	return false
}
