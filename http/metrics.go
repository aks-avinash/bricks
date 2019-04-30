package http

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	kb = 1024
	mb = kb * kb
)

var (
	paceHTTPInFlightGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pace_http_in_flight_requests",
		Help: "A gauge of requests currently being served by the wrapped handler.",
	})

	paceHTTPCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "pace_api_http_request_total",
			Help: "A counter for requests to the wrapped handler.",
		},
		[]string{"code", "method", "source", "path"},
	)

	// Duration is labeled by the request method and source, and response code.
	// It uses custom buckets based on the expected request duration.
	paceHTTPDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "pace_api_http_request_duration_seconds",
			Help:    "A histogram of request durations.",
			Buckets: []float64{0.01, 0.05, 0.1, 0.3, 0.6, 1, 2.5, 5, 10, 60},
		},
		[]string{"code", "method", "source", "path"},
	)

	// Size is labeled by the request method, path and source, response code, and type.
	// The type label distinguishes between request and response size.
	// It uses custom buckets based on the expected response size.
	paceHTTPSize = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "pace_api_http_size_bytes",
			Help: "A histogram of request and response body sizes.",
			Buckets: []float64{
				100,
				kb, 10 * kb, 100 * kb,
				1 * mb, 5 * mb, 10 * mb, 100 * mb,
			},
		},
		[]string{"code", "method", "source", "path", "type"},
	)
)

func init() {
	// Register all of the metrics in the standard registry.
	prometheus.MustRegister(paceHTTPInFlightGauge, paceHTTPCounter, paceHTTPDuration, paceHTTPSize)
}

func metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paceHTTPInFlightGauge.Inc()
		defer paceHTTPInFlightGauge.Dec()

		startTime := time.Now()
		srw := statusWriter{ResponseWriter: w}
		next.ServeHTTP(&srw, r)
		dur := float64(time.Since(startTime)) / float64(time.Millisecond)

		labels := prometheus.Labels{
			"code":   strconv.Itoa(srw.status),
			"method": r.Method,
			"source": filterRequestSource(r.Header.Get("Request-Source")),
			"path":   "",
		}
		if route := mux.CurrentRoute(r); route != nil {
			if path, err := route.GetPathTemplate(); err == nil {
				labels["path"] = path
			}
		}

		paceHTTPCounter.With(labels).Inc()
		paceHTTPDuration.With(labels).Observe(dur)
		labels["type"] = "req"
		paceHTTPSize.With(labels).Observe(float64(r.ContentLength)) // request size
		labels["type"] = "resp"
		paceHTTPSize.With(labels).Observe(float64(srw.length)) // response size
	})
}

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}

func filterRequestSource(source string) string {
	switch source {
	case "uptime", "kubernetes", "nginx", "livetest":
		return source
	}
	return ""
}
