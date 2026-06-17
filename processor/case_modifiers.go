package processor

import (
	"regexp"
	"strconv"
	"strings"
)

var caseModRe = regexp.MustCompile(`\((up|low|cap)(?:,\s*(\d+))?\)`)

// caseApply transforms a single word according to the given mode.
// Private to this file — not used anywhere else.
func caseApply(word, mode string) string {
	switch mode {
	case "up":
		return strings.ToUpper(word)
	case "low":
		return strings.ToLower(word)
	case "cap":
		if len(word) == 0 {
			return word
		}
		return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	}
	return word
}

// HandleCaseModifiers finds every (up), (low), (cap) — with an optional word
// count — and transforms the preceding word(s). Works directly on a string.
func HandleCaseModifiers(text string) string {
	for {
		loc := caseModRe.FindStringSubmatchIndex(text)
		if loc == nil {
			break
		}

		mode := text[loc[2]:loc[3]]
		count := 1
		if loc[4] != -1 {
			if n, err := strconv.Atoi(text[loc[4]:loc[5]]); err == nil && n > 0 {
				count = n
			}
		}

		before := strings.TrimRight(text[:loc[0]], " ")
		after := strings.TrimLeft(text[loc[1]:], " ")

		words := strings.Fields(before)
		start := len(words) - count
		if start < 0 {
			start = 0
		}
		for i := start; i < len(words); i++ {
			words[i] = caseApply(words[i], mode)
		}

		text = strings.Join(words, " ")
		if after != "" {
			text += " " + after
		}
	}
	return text
}
