package analyzer

import (
	"strings"
	"unicode"
)

func lowercase_make(text string) string {
	new := []rune(text)
	for i, r := range text {
		if i == 0 {
			new[i] = unicode.ToLower(r)
			break
		}
	}
	return string(new)
}

func symbols_clean(text string) string {
	result := make([]rune, 0, len(text))
	for _, r := range text {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != ' ' {
			result = append(result, ' ')
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func sensetive_clean(text string, cfg *Config) string {
	patterns := cfg.get_sensitive_patterns()
	
	patternMap := make(map[string]bool)
	for _, p := range patterns {
		patternMap[strings.ToLower(p)] = true
	}
	words := strings.Fields(text)
	result := make([]string, 0, len(words))
	
	for _, word := range words {
		shouldRedact := false
		wordLower := strings.ToLower(word)
		cleanWord := strings.TrimFunc(wordLower, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r)
		})
		
		for pattern := range patternMap {
			if strings.Contains(cleanWord, pattern) {
				shouldRedact = true
				break
			}
		}
		if shouldRedact {
			result = append(result, "deleted")
		} else {
			result = append(result, word)
		}
	}
	
	return strings.Join(result, " ")
}

func clean_no_english(text string) string {
	result := make([]rune, 0, len(text))

	for _, r := range text {
		if unicode.IsLetter(r) && unicode.In(r, unicode.Latin) {
			result = append(result, r)
		} else if !unicode.IsLetter(r) {
			result = append(result, r)
		} else {
			result = append(result, ' ')
		}
	}

	return string(result)
}
