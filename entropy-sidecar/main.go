package main

import (
	"entropy-sidecar/cache"
	"entropy-sidecar/handlers"
	"entropy-sidecar/metrics"
	"entropy-sidecar/proxy"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	redisClient := cache.NewRedisClient()
	metricsStore := metrics.NewMetrics()

	cacheProxy := proxy.NewCacheProxy(
		"http://localhost:9090",
		redisClient,
		15*time.Second,
		metricsStore,
	)

	mux.HandleFunc("/sidecar/health", handlers.SidecarHealthHandler)
	mux.HandleFunc("/sidecar/metrics", handlers.MetricsHandler(metricsStore))
	mux.Handle("/", cacheProxy)

	log.Println("Entropy Sidecar is running on port 8080")
	log.Println("Proxying requests to Desk Mess Service on port 9090")
	log.Println("Redis cache is running on localhost:6379")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}