package domain

type RateLimited interface {
	RetryAfter() int64
}

type rateLimited struct {
	retryAfter int64
}

func (r *rateLimited) RetryAfter() int64 {
	return r.retryAfter
}
