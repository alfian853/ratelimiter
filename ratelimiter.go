package ratelimiter

type RateLimiter interface {
	Execute(string, func()) bool
}
