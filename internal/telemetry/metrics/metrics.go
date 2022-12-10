package metrics

import (
	"errors"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	Namespace = "stallion_blackbox_exporter"
	Subsystem = "client"
)

// Metrics has all the client metrics.
type Metrics struct {
	ConnectionErrors prometheus.Counter
	PublishErrors    prometheus.Counter
	ConsumeErrors    prometheus.Counter
	ResponseTime     prometheus.Histogram
}

// newCounter generator.
func newCounter(counterOpts prometheus.CounterOpts) prometheus.Counter {
	ev := prometheus.NewCounter(counterOpts)

	if err := prometheus.Register(ev); err != nil {
		var are prometheus.AlreadyRegisteredError
		if ok := errors.As(err, &are); ok {
			ev, ok = are.ExistingCollector.(prometheus.Counter)
			if !ok {
				panic("different metric type registration")
			}
		} else {
			panic(err)
		}
	}

	return ev
}

// newHistogram generator.
func newHistogram(histogramOpts prometheus.HistogramOpts) prometheus.Histogram {
	ev := prometheus.NewHistogram(histogramOpts)

	if err := prometheus.Register(ev); err != nil {
		var are prometheus.AlreadyRegisteredError
		if ok := errors.As(err, &are); ok {
			ev, ok = are.ExistingCollector.(prometheus.Histogram)
			if !ok {
				panic("different metric type registration")
			}
		} else {
			panic(err)
		}
	}

	return ev
}

// NewMetrics creates a new metrics type for black box status.
func NewMetrics() Metrics {
	return Metrics{
		ConnectionErrors: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "connection_errors_total",
			Help:        "total number of connection errors",
			ConstLabels: nil,
		}),
		PublishErrors: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "publish_errors_total",
			Help:        "total number of publish errors",
			ConstLabels: nil,
		}),
		ConsumeErrors: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "consume_errors_total",
			Help:        "total number of consume errors",
			ConstLabels: nil,
		}),
		ResponseTime: newHistogram(prometheus.HistogramOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "response_duration_seconds",
			Help:        "from ping to pong duration in seconds",
			ConstLabels: nil,
			Buckets:     prometheus.DefBuckets,
		}),
	}
}
