package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	TotalRequests *prometheus.CounterVec
	HttpDuration  *prometheus.HistogramVec
	ErrorRequests *prometheus.CounterVec
}

var (
	instance *Metrics
	once     sync.Once
)

func NewMetrics(reg prometheus.Registerer) *Metrics {
	once.Do(func() {
		const namePrefix = "pushart_"
		instance = &Metrics{
			TotalRequests: prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Name: namePrefix + "requests_total",
					Help: "Number of get requests.",
				},
				[]string{"path", "method", "code", "handler"},
			),
			HttpDuration: promauto.NewHistogramVec(
				prometheus.HistogramOpts{
					Name: namePrefix + "response_time_seconds",
					Help: "Duration of HTTP requests.",
				},
				[]string{"path", "method", "code", "handler"},
			),
			ErrorRequests: prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Name: namePrefix + "errors_total",
					Help: "Number of error responses.",
				},
				[]string{"path", "method", "code", "handler"},
			),
		}
		reg.MustRegister(instance.TotalRequests, instance.HttpDuration, instance.ErrorRequests)
	})
	return instance
}

// GetMetrics возвращает уже созданные метрики
func GetMetrics() *Metrics {
	return instance
}
