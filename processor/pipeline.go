package processor

func Process(text string) string {
	return applyTransformations(text)
}

// applyTransformations runs all stages in order.
// order matters: each stage must complete before the next one begins.
func applyTransformations(text string) string {
	text = HandleHexBin(text)
	text = HandleCaseModifiers(text)
	text = FixPunctuation(text)
	text = FixArticles(text)
	return text
}
