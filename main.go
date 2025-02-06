
package main

import (
	"fmt"
	"strings"
)

// Simple sentiment analysis function
func analyzeSentiment(text string) string {
	text = strings.ToLower(text)
	positiveKeywords := []string{"good", "great", "excellent", "happy", "love"}
	negativeKeywords := []string{"bad", "terrible", "horrible", "sad", "hate"}

	positiveScore := 0
	negativeScore := 0
	
	words := strings.Fields(text)
	for _, word := range words {
		for _, pk := range positiveKeywords {
			if strings.Contains(word, pk) {
				positiveScore++
			}
		}
		for _, nk := range negativeKeywords {
			if strings.Contains(word, nk) {
				negativeScore++
			}
		}
	}

	if positiveScore > negativeScore {
		return "Positive"
	} else if negativeScore > positiveScore {
		return "Negative"
	} else {
		return "Neutral"
	}
}

func main() {
	phrases := []string{
		"This is a great movie, I love it!",
		"The service was terrible and I hate it.",
		"It was an okay experience.",
		"I am happy with the excellent results.",
	}

	for _, phrase := range phrases {
		fmt.Printf("Phrase: "%s" -> Sentiment: %s
", phrase, analyzeSentiment(phrase))
	}
}
