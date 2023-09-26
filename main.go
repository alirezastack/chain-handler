package main

import (
	chain "github.com/alirezastack/chain-handler/chains"
	"github.com/rs/zerolog/log"
)

func main() {
	// initialize last chain handler first
	httpHandler := chain.HttpHandler{}

	// initialize other chain handlers
	cacheHandler := chain.CacheHandler{}
	cacheHandler.SetNext(&httpHandler)

	providerName := cacheHandler.Execute("sample-provider-id")

	log.Info().Msgf("provider name is: %s", providerName)
}
