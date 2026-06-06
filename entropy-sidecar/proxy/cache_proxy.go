package proxy

import (
	"context"
	"entropy-sidecar/metrics"
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
	metrics  *metrics.Metrics
}

func NewCacheProxy(
	target string,
	redisClient *redis.Client,
	cacheTTL time.Duration,
	metricsStore *metrics.Metrics,
) *CacheProxy {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)

	return &CacheProxy{
		proxy:    reverseProxy,
		redis:    redisClient,
		cacheTTL: cacheTTL,
		metrics:  metricsStore,
	}
}

func (c *CacheProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	if r.Method != http.MethodGet {
		recorder := newResponseRecorder(w)

		c.proxy.ServeHTTP(recorder, r)

		logRequest(r, recorder.statusCode, "BYPASS", time.Since(start))
		return
	}

	cacheKey := buildCacheKey(r)

	cachedBody, err := c.redis.Get(context.Background(), cacheKey).Result()
	if err == nil {
		c.metrics.RecordCacheHit()

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Cache", "HIT")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(cachedBody))

		logRequest(r, http.StatusOK, "HIT", time.Since(start))
		return
	}

	if err != redis.Nil {
		log.Printf("Redis get error: %v", err)
	}

	c.metrics.RecordCacheMiss()

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

	logRequest(r, recorder.statusCode, "MISS", time.Since(start))
}

func buildCacheKey(r *http.Request) string {
	return r.Method + ":" + r.URL.Path
}

func logRequest(r *http.Request, statusCode int, cacheStatus string, duration time.Duration) {
	log.Printf(
		"method=%s path=%s status=%d cache=%s duration=%s",
		r.Method,
		r.URL.Path,
		statusCode,
		cacheStatus,
		duration,
	)
}