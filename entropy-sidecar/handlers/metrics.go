package handlers

import (
	"entropy-sidecar/metrics"
	"net/http"
)

func MetricsHandler(metricsStore *metrics.Metrics) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{
				"error":   "method_not_allowed",
				"message": "Only GET method is allowed",
			})
			return
		}

		WriteJSON(w, http.StatusOK, metricsStore.Snapshot())
	}
}