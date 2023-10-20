package ratelimiter

import (
	"fmt"
)

type SimpleRateLimiter struct {
	limitCounterMap map[string]chan bool
	limitPerKey     int
}

func NewSimpleRateLimiter(limitPerKey int) *SimpleRateLimiter {
	return &SimpleRateLimiter{limitPerKey: limitPerKey, limitCounterMap: map[string]chan bool{}}
}

func (r *SimpleRateLimiter) Execute(key string, executable func()) bool {

	fmt.Println("handle request")
	counterCh, ok := r.limitCounterMap[key]
	if !ok {
		counterCh = make(chan bool, r.limitPerKey)
		r.limitCounterMap[key] = counterCh
	}

	if len(counterCh) >= r.limitPerKey {
		return false
	}
	counterCh <- true
	executable()
	<-counterCh
	return true
}
