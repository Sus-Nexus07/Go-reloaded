package processor

import (
	"fmt"
	"regexp"
	"strconv"
)

var hexBinRe = regexp.MustCompile(`(\S+)\s+\((hex|bin)\)`)

// HandleHexBin finds every (hex) or (bin) marker and replaces the preceding
// word with its decimal value. Works directly on a string — no other function needed.
func HandleHexBin(text string) string {
	return hexBinRe.ReplaceAllStringFunc(text, func(match string) string {
		parts := hexBinRe.FindStringSubmatch(match)
		base := 16
		if parts[2] == "bin" {
			base = 2
		}
		val, err := strconv.ParseInt(parts[1], base, 64)
		if err != nil {
			return match // leave invalid input untouched
		}
		return fmt.Sprintf("%d", val)
	})
}
