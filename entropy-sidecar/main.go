package main

import (
	"entropy-sidecar/cache"
	"entropy-sidecar/handlers"
	"entropy-sidecar/metrics"
	"entropy-sidecar/proxy"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()

	targetServiceURL := os.Getenv("TARGET_SERVICE_URL")
	if targetServiceURL == "" {
		targetServiceURL = "http://localhost:9090"
	}

	redisClient := cache.NewRedisClient()
	metricsStore := metrics.NewMetrics()

	cacheProxy := proxy.NewCacheProxy(
		targetServiceURL,
		redisClient,
		15*time.Second,
		metricsStore,
	)

	mux.HandleFunc("/sidecar/health", handlers.SidecarHealthHandler)
	mux.HandleFunc("/sidecar/metrics", handlers.MetricsHandler(metricsStore))
	mux.Handle("/", cacheProxy)

	log.Println("Entropy Sidecar is running on port 8080")
	log.Printf("Proxying requests to Desk Mess Service at %s", targetServiceURL)
	log.Printf("Redis cache is running on %s", os.Getenv("REDIS_ADDR"))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}