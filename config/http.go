package config

import (
	"net/http"

	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
)

func Http() *http.Client {
	cacheTransport := httpcache.NewTransport(diskcache.New(".cache"))

	client := &http.Client{
		Transport: cacheTransport,
	}

	return client
}
