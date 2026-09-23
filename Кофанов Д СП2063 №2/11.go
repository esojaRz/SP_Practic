package main

import (
	"fmt"
	"strings"
)
type TextStatistics struct {
	CharCount int
	WordCount int
	SentenceCount int
}
func textStats(text string) TextStatistics {
	words := strings.Fields(text)
	sentences := 0
	for _, char := range text {
		if char == '.' || char == '!' || char == '?' {
			sentences++
		}
	}
	return TextStatistics{
		CharCount:     len([]rune(text)),
		WordCount:     len(words),
		SentenceCount: sentences,
	}
}
func main() {
	text := "Привет! Как твои дела? Надеюсь, все отлично."
	stats := textStats(text)
	fmt.Printf("Символов: %d, Слов: %d, Предложений: %d\n", stats.CharCount, stats.WordCount, stats.SentenceCount)
}