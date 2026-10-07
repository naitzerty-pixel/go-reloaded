package main

import "strings"

// Process enchaîne toutes les étapes de transformation.
func Process(text string) string {
	words := strings.Fields(text)
	words = ApplyMarkers(words) // hex, bin, up, low, cap
	text = strings.Join(words, " ")
	text = FixPunctuation(text)
	text = FixQuotes(text)
	text = FixArticles(text)
	return text
}
