package stats

import (
	"time"

	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
)

// NewLogReporter returns a tally reporter that logs to the provided logger at debug level. This is useful for
// local development where the usual sinks are not available or refreshing the prometheus endpoint is not desirable.
func NewLogReporter(logger *zap.Logger) tally.StatsReporter {
	return &logReporter{logger: logger}
}

type logReporter struct {
	logger *zap.Logger
}

// Capabilities interface.
func (r *logReporter) Reporting() bool {
	return true
}

func (r *logReporter) Tagging() bool {
	return true
}

func (r *logReporter) Capabilities() tally.Capabilities {
	return r
}

// Reporter interface.
func (r *logReporter) Flush() {
	// Silence.
}

func (r *logReporter) ReportCounter(name string, tags map[string]string, value int64) {
	r.logger.Debug("counter",
		zap.String("name", name),
		zap.Int64("value", value),
		zap.Any("tags", tags),
		zap.String("type", "counter"),
	)
}

func (r *logReporter) ReportGauge(name string, tags map[string]string, value float64) {
	r.logger.Debug("gauge",
		zap.String("name", name),
		zap.Float64("value", value),
		zap.Any("tags", tags),
		zap.String("type", "gauge"),
	)
}

func (r *logReporter) ReportTimer(name string, tags map[string]string, interval time.Duration) {
	r.logger.Debug("timer",
		zap.String("name", name),
		zap.Duration("value", interval),
		zap.Any("tags", tags),
		zap.String("type", "timer"),
	)
}

func (r *logReporter) ReportHistogramValueSamples(
	name string,
	tags map[string]string,
	buckets tally.Buckets,
	bucketLowerBound,
	bucketUpperBound float64,
	samples int64,
) {
	r.logger.Debug("histogram",
		zap.String("name", name),
		zap.Float64s("buckets", buckets.AsValues()),
		zap.Float64("bucketLowerBound", bucketLowerBound),
		zap.Float64("bucketUpperBound", bucketUpperBound),
		zap.Int64("samples", samples),
		zap.Any("tags", tags),
		zap.String("type", "valueHistogram"),
	)
}

func (r *logReporter) ReportHistogramDurationSamples(
	name string,
	tags map[string]string,
	buckets tally.Buckets,
	bucketLowerBound,
	bucketUpperBound time.Duration,
	samples int64,
) {
	r.logger.Debug("histogram",
		zap.String("name", name),
		zap.Durations("buckets", buckets.AsDurations()),
		zap.Duration("bucketLowerBound", bucketLowerBound),
		zap.Duration("bucketUpperBound", bucketUpperBound),
		zap.Int64("samples", samples),
		zap.Any("tags", tags),
		zap.String("type", "durationHistogram"),
	)
}
