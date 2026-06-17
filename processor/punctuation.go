package processor

import (
	"regexp"
	"strings"
)

var (
	puncBefore = regexp.MustCompile(`\s+([.,!?:;]+)`)           // spaces before punctuation
	puncAfter  = regexp.MustCompile(`([.,!?:;]+)([^\s.,!?:;])`) // missing space after punctuation
	quoteOpen  = regexp.MustCompile(`'\s+`)                     // space after opening quote
	quoteClose = regexp.MustCompile(`\s+'`)                     // space before closing quote
)

// FixPunctuation corrects spacing around punctuation marks and single quotes.
// Works directly on a string — no other function needed.
func FixPunctuation(text string) string {
	text = puncBefore.ReplaceAllString(text, "$1")   // "word , next" → "word, next"
	text = puncAfter.ReplaceAllString(text, "$1 $2") // "word,next"   → "word, next"
	text = quoteOpen.ReplaceAllString(text, "'")     // "' word"      → "'word"
	text = quoteClose.ReplaceAllString(text, "'")    // "word '"      → "word'"
	return strings.Join(strings.Fields(text), " ")   // normalise any leftover whitespace
}
