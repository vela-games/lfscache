package exporter

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type LFSCacheCollector struct {
	CacheLatency metrics.Histogram
	Source       metrics.Counter
	TransferRate metrics.Gauge
}

func NewCollector() *LFSCacheCollector {
	return &LFSCacheCollector{
		Source: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "lfscache",
			Name:      "cache_source",
			Help:      "Cache source",
		}, []string{"type"}),
		TransferRate: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: "lfscache",
			Name:      "cache_transfer_rate_bytes",
			Help:      "Transfer rate for outgoing requests",
		}, []string{}),
		CacheLatency: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: "lfscache",
			Name:      "cache_latency",
			Help:      "Latency for cache retrieval",
		}, []string{"source"}),
	}
}
