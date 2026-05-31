package proxy

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheProxy struct {
	proxy    *httputil.ReverseProxy
	redis    *redis.Client
	cacheTTL time.Duration
}

func NewCacheProxy(target string, redisClient *redis.Client, cacheTTL time.Duration) *CacheProxy {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)

	return &CacheProxy{
		proxy:    reverseProxy,
		redis:    redisClient,
		cacheTTL: cacheTTL,
	}
}

func (c *CacheProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		c.proxy.ServeHTTP(w, r)
		return
	}

	cacheKey := buildCacheKey(r)

	cachedBody, err := c.redis.Get(context.Background(), cacheKey).Result()
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Cache", "HIT")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(cachedBody))
		return
	}

	if err != redis.Nil {
		log.Printf("Redis get error: %v", err)
	}

	recorder := newResponseRecorder(w)
	recorder.Header().Set("X-Cache", "MISS")

	c.proxy.ServeHTTP(recorder, r)

	if recorder.statusCode == http.StatusOK {
		err := c.redis.Set(
			context.Background(),
			cacheKey,
			recorder.body.String(),
			c.cacheTTL,
		).Err()

		if err != nil {
			log.Printf("Redis set error: %v", err)
		}
	}
}

func buildCacheKey(r *http.Request) string {
	return r.Method + ":" + r.URL.Path
}