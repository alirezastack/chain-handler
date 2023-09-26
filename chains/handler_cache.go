package chain

import (
	"context"
	"time"
)

type CacheHandler struct {
	next BaseHandler[string]
}

// Execute let's assume cacheHandler hits redis cache, if not found in redis cache it should call remote http service (next handler)
func (handler *CacheHandler) Execute(id string) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// BELOW LINE IS NOT IMPLEMENTED
	value, _ := redisClient.Get(ctx, id).Result()
	if value == "" && handler.next != nil {
		return handler.next.Execute(id)
	}

	return value
}

func (handler *CacheHandler) SetNext(next BaseHandler[string]) {
	handler.next = next
}
