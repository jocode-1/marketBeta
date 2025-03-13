package config

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
)

func InitMonitoring() {
	prometheus.MustRegister(httpRequestsTotal, httpRequestDuration)
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":9090", nil)
}

// RequestMetricsMiddleware tracks request count & duration
func RequestMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next() // Process request

		duration := time.Since(start).Seconds()
		statusCode := c.Writer.Status()

		httpRequestsTotal.WithLabelValues(c.Request.Method, c.FullPath(), http.StatusText(statusCode)).Inc()
		httpRequestDuration.WithLabelValues(c.Request.Method, c.FullPath()).Observe(duration)
	}
}
