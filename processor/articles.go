package processor

import (
	"regexp"
	"strings"
)

var articleRe = regexp.MustCompile(`(?i)\b(an?)\b\s+(\S+)`)

// FixArticles corrects "a"/"an" before vowel sounds (a, e, i, o, u, h).
// Preserves the original capitalisation of the article.
// Works directly on a string — no other function needed.
func FixArticles(text string) string {
	return articleRe.ReplaceAllStringFunc(text, func(match string) string {
		parts := articleRe.FindStringSubmatch(match)
		article := parts[1]
		nextWord := parts[2]

		// Strip leading punctuation to reach the first real letter.
		clean := strings.TrimLeft(nextWord, ".,!?:;'\"()")
		if len(clean) == 0 {
			return match
		}

		needsAn := strings.ContainsRune("aeiouhAEIOUH", rune(clean[0]))
		lower := strings.ToLower(article)
		isCapital := article[0] >= 'A' && article[0] <= 'Z'

		switch {
		case needsAn && lower == "a": // "a apple" → "an apple"
			if isCapital {
				return "An " + nextWord
			}
			return "an " + nextWord
		case !needsAn && lower == "an": // "an dog" → "a dog"
			if isCapital {
				return "A " + nextWord
			}
			return "a " + nextWord
		}
		return match
	})
}
