package analyzer

import (
	"go/ast"
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
		if !unicode.IsLetter(letter) && !unicode.IsDigit(letter) && letter != ' ' {
			return false
		}
	}
	return true
}

func english_check(text string) bool {
	for _, letter := range text {
		if unicode.IsLetter(letter) && !unicode.In(letter, unicode.Latin) {
			return false
		}
	}
	return true
}
func sensitive_check(text string, cfg *Config) bool {
	patterns := cfg.get_sensitive_patterns()
	lowerText := strings.ToLower(text)

	for _, pattern := range patterns {
		if strings.Contains(lowerText, pattern) {
			return false
		}
	}
	return true
}

func checkRules(pass *analysis.Pass, lit *ast.BasicLit, text string, cfg *Config) {
	if text == "" {
		return
	}
	if cfg.Rules.CheckSensitiveData && !sensitive_check(text, cfg) {
		fixed_text := sensetive_clean(text, cfg)

		pass.Report(analysis.Diagnostic{
			Pos:     lit.Pos(),
			End:     lit.End(),
			Message: "Log messages must not contain sensitive data\n",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "redact sensitive data",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     lit.Pos(),
							End:     lit.End(),
							NewText: []byte(`"` + fixed_text + `"`),
						},
					},
				},
			},
		})
	}
	if cfg.Rules.CheckFirstLetter && !lowercase_check(text) {
		fixed_text := lowercase_make(text)

		pass.Report(analysis.Diagnostic{
			Pos:     lit.Pos(),
			End:     lit.End(),
			Message: "Log messages must start with a lowercase letter\n",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "make first letter lowercase",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     lit.Pos(),
							End:     lit.End(),
							NewText: []byte(`"` + fixed_text + `"`),
						},
					},
				},
			},
		})
	}

	if cfg.Rules.CheckEnglish && !english_check(text) {
		fixed_text := clean_no_english(text)

		pass.Report(analysis.Diagnostic{
			Pos:     lit.Pos(),
			End:     lit.End(),
			Message: "Log messages must be in English\n",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "replace non-english letters",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     lit.Pos(),
							End:     lit.End(),
							NewText: []byte(`"` + fixed_text + `"`),
						},
					},
				},
			},
		})
	}

	if cfg.Rules.CheckSpecialChars && !symbols_check(text) {
		fixed_text := symbols_clean(text)

		pass.Report(analysis.Diagnostic{
			Pos:     lit.Pos(),
			End:     lit.End(),
			Message: "Log messages must not contain special characters or emojis\n",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "remove special characters",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     lit.Pos(),
							End:     lit.End(),
							NewText: []byte(`"` + fixed_text + `"`),
						},
					},
				},
			},
		})
	}

	
}
