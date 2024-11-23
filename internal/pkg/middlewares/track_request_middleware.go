package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"gopkg.in/alexcesaro/statsd.v2"
)

type Timing struct {
	Count    int
	Duration time.Duration
}

type ctxTimings struct {
	sync.Mutex
	Data map[string]*Timing
}

type key int

const timingsKey key = 1

type TimingMiddleware struct {
	sync.Mutex
	StatsReciever *statsd.Client
	Metrics       map[string]int
}

func (tm *TimingMiddleware) TrackRequestTimings(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx,
			timingsKey,
			&ctxTimings{
				Data: make(map[string]*Timing),
			})
		defer tm.logContextTimings(ctx, r.URL.Path, time.Now())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (tm *TimingMiddleware) logContextTimings(ctx context.Context, path string, start time.Time) {
	// получаем тайминги из контекста
	// поскольку там пустой интерфейс, то нам надо преобразовать к нужному типу
	timings, ok := ctx.Value(timingsKey).(*ctxTimings)
	if !ok {
		return
	}
	totalReal := time.Since(start)

	path = strings.Replace(path, "/", "-", -1)

	prefix := "requests." + path + "."

	var total time.Duration
	for timing, value := range timings.Data {
		metric := prefix + "timings." + timing
		tm.StatsReciever.Increment(metric)
		tm.StatsReciever.Timing(metric+"_time", uint64(value.Duration/time.Millisecond))
		total += value.Duration
	}

	tm.StatsReciever.Increment(prefix + "hits")
	tm.StatsReciever.Timing(prefix+"tracked", uint64(totalReal/time.Millisecond))
	tm.StatsReciever.Timing(prefix+"real_time", uint64(total/time.Millisecond))
	log.Println(prefix+"real_time", uint64(total/time.Millisecond))
}
