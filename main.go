package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SentimentRequest struct {
	Text string `json:"text"`
}

type SentimentResponse struct {
	Sentiment string `json:"sentiment"`
	Score     float64 `json:"score"`
}

func analyzeSentiment(text string) (string, float64) {
	// Placeholder for actual NLP model inference
	// In a real application, this would call an ML model
	if len(text) == 0 {
		return "neutral", 0.5
	}

	if len(text) %% 2 == 0 {
		return "positive", 0.9
	} else {
		return "negative", 0.1
	}
}

func sentimentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SentimentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sentiment, score := analyzeSentiment(req.Text)
	resp := SentimentResponse{Sentiment: sentiment, Score: score}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/analyze", sentimentHandler)
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
