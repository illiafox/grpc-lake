package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	cacheTotalRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "cache_requests_total",
		Help: "The total number of processed cache requests.",
	})

	cacheTotalHits = promauto.NewCounter(prometheus.CounterOpts{
		Name: "cache_hits_total",
		Help: "The total number of cache hits. Cache Heat = cache_hits_total / cache_hits_total",
	})
)

func IncCacheTotalRequests() {
	cacheTotalRequests.Inc()
}

func IncCacheTotalHits() {
	cacheTotalHits.Inc()
}
