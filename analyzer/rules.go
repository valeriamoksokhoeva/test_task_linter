package analyzer

import (
	"go/token"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func lowercase_check(text string) bool {
	if text == "" {
		return true
	}
	for i, r := range text {
		if i == 0 {
			if unicode.IsLetter(r) && !unicode.IsUpper(r) {
				return true
			}
			break
		}
	}
	return false
}

func symbols_check(text string) bool {
	for _, letter := range text {
		if !(unicode.IsLetter(letter) || unicode.IsDigit(letter) || letter == ' '){
			return false
		}
	}
	return true
}

func english_check(text string) bool {
	for _, letter := range text {
		if !(unicode.In(letter, unicode.Latin) || letter == ' ' || !unicode.IsLetter(letter)){
			return false
		}
	}
	return true
}

func sensitive_check(text string) bool {
	lower_text := strings.ToLower(text)
	if strings.Contains(lower_text, "key") || strings.Contains(lower_text, "token") || strings.Contains(lower_text, "password") || strings.Contains(lower_text, "secret") {
		return false
	}
	return true
}

func CheckRules(pass *analysis.Pass, pos token.Pos, text string) {
	if !lowercase_check(text) {
		pass.Reportf(pos, "Log messages must start with a lowercase letter\n")
	}
	if !english_check(text) {
		pass.Reportf(pos, "Log messages must be in English\n")
	}
	if !symbols_check(text) {
		pass.Reportf(pos, "Log messages must not contain special characters or emojis\n")
	}
	if !sensitive_check(text) {
		pass.Reportf(pos, "Log messages must not contain sensitive data\n")
	}
}