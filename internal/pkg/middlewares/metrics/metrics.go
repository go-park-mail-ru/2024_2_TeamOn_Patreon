package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
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
				[]string{"path", "method", "code"},
			),
			HttpDuration: prometheus.NewHistogramVec(
				prometheus.HistogramOpts{
					Name: namePrefix + "response_time_seconds",
					Help: "Duration of HTTP requests.",
				},
				[]string{"path", "method", "code"},
			),
			ErrorRequests: prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Name: namePrefix + "errors_total",
					Help: "Number of error responses.",
				},
				[]string{"path", "method", "code"},
			),
		}
		reg.MustRegister(instance.TotalRequests)
		reg.MustRegister(instance.HttpDuration)
		reg.MustRegister(instance.ErrorRequests)
	})
	return instance
}

// GetMetrics возвращает уже созданные метрики
func GetMetrics() *Metrics {
	return instance
}
