package middlewares

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/middlewares/metrics"
)

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rec, r)

		// Обновляем метрики
		m := metrics.GetMetrics()
		m.TotalRequests.WithLabelValues(r.URL.Path, r.Method, strconv.Itoa(rec.statusCode)).Inc()
		m.HttpDuration.WithLabelValues(r.URL.Path, r.Method, strconv.Itoa(rec.statusCode)).Observe(time.Since(start).Seconds())

		// Инкремент счетчика ошибок
		if rec.statusCode >= 400 {
			m.ErrorRequests.WithLabelValues(r.URL.Path, r.Method, strconv.Itoa(rec.statusCode)).Inc()
		}
	})
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}
