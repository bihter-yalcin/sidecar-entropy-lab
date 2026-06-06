package metrics

import "sync"

type Metrics struct {
	mu              sync.Mutex
	TotalRequests   int     `json:"total_requests"`
	CacheHits       int     `json:"cache_hits"`
	CacheMisses     int     `json:"cache_misses"`
	ProxiedRequests int     `json:"proxied_requests"`
	HitRatio        float64 `json:"hit_ratio"`
}

func NewMetrics() *Metrics {
	return &Metrics{}
}

func (m *Metrics) RecordCacheHit() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.TotalRequests++
	m.CacheHits++
	m.updateHitRatio()
}

func (m *Metrics) RecordCacheMiss() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.TotalRequests++
	m.CacheMisses++
	m.ProxiedRequests++
	m.updateHitRatio()
}

func (m *Metrics) Snapshot() Metrics {
	m.mu.Lock()
	defer m.mu.Unlock()

	return Metrics{
		TotalRequests:   m.TotalRequests,
		CacheHits:       m.CacheHits,
		CacheMisses:     m.CacheMisses,
		ProxiedRequests: m.ProxiedRequests,
		HitRatio:        m.HitRatio,
	}
}

func (m *Metrics) updateHitRatio() {
	if m.TotalRequests == 0 {
		m.HitRatio = 0
		return
	}

	m.HitRatio = float64(m.CacheHits) / float64(m.TotalRequests)
}