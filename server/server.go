package main

import (
	"fmt"
	"net/http"
	"ratelimiter"
	"time"
)

var rateLimiter ratelimiter.RateLimiter = ratelimiter.NewSimpleRateLimiter(10)

func HandleLimitByURLAndIdentifier(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("user-id")
	if userID == "" {
		w.WriteHeader(400)
		return
	}
	key := "HandleLimitByURLAndIdentifier:" + userID

	isExecuted := rateLimiter.Execute(key, func() {
		time.Sleep(1 * time.Second)
		w.WriteHeader(203)
	})

	if !isExecuted {
		w.WriteHeader(429)
	}
}

func main() {
	http.HandleFunc("/limit-by-url-and-identifier", HandleLimitByURLAndIdentifier)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
