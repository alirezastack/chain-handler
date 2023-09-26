package chain

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type HttpHandler struct {
	next BaseHandler[string]
}

func (handler *HttpHandler) Execute(id string) string {
	url := fmt.Sprintf("https://third-party-url.com/%s", id)
	providerRS, err := http.Get(url)
	if err != nil {
		log.Warn().Msgf("Could not get provider info from 3rd party service: %s", err)
		if handler.next != nil {
			return handler.next.Execute(id)
		}
		return ""
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// suppose providerRS.Status is JSON API response (provider name)
	providerName := providerRS.Status

	// NOT IMPLEMENTED
	err = redisClient.Set(ctx, id, providerName, 0).Err()
	if err != nil {
		log.Error().Msgf("provider %s failed to be set in cache service: %s", providerRS, err)
	}

	return providerName
}

func (handler *HttpHandler) SetNext(next BaseHandler[string]) {
	handler.next = next
}
