package config

import (
	"net/http"
	"time"

	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
)

func Http() *http.Client {
	cacheTransport := httpcache.NewTransport(diskcache.New(".cache"))

	client := &http.Client{
		Transport: cacheTransport,
		Timeout:   5 * time.Minute,
	}

	return client
}
