package telegraf

import (
	"time"
)

// Accumulator 允许往处理流中添加metrics
type Accumulator interface {
	// AddFields adds a metric to the accumulator with the given measurement
	// name, fields, and tags (and timestamp). If a timestamp is not provided,
	// then the accumulator sets it to "now".
	AddFields(measurement string,
		fields map[string]interface{},
		tags map[string]string,
		t ...time.Time)

	// AddGauge is the same as AddFields, but will add the metric as a "Gauge" type
	AddGauge(measurement string,
		fields map[string]interface{},
		tags map[string]string,
		t ...time.Time)

	// AddCounter is the same as AddFields, but will add the metric as a "Counter" type
	AddCounter(measurement string,
		fields map[string]interface{},
		tags map[string]string,
		t ...time.Time)

	// AddSummary is the same as AddFields, but will add the metric as a "Summary" type
	AddSummary(measurement string,
		fields map[string]interface{},
		tags map[string]string,
		t ...time.Time)

	// AddHistogram is the same as AddFields, but will add the metric as a "Histogram" type
	AddHistogram(measurement string,
		fields map[string]interface{},
		tags map[string]string,
		t ...time.Time)

	// AddMetric adds an metric to the accumulator.
	AddMetric(Metric)

	// SetPrecision sets the timestamp rounding precision.  All metrics addeds
	// added to the accumulator will have their timestamp rounded to the
	// nearest multiple of precision.
	SetPrecision(precision time.Duration)

	// Report an error.
	AddError(err error)

	// Upgrade to a TrackingAccumulator with space for maxTracked
	// metrics/batches.
	WithTracking(maxTracked int) TrackingAccumulator
}

// TrackingID 唯一标识一个追踪的metric组
type TrackingID uint64

// DeliveryInfo 返回交付后的结果
type DeliveryInfo interface {
	// ID is the TrackingID
	ID() TrackingID

	// Delivered returns true if the metric was processed successfully.
	Delivered() bool
}

// TrackingAccumulator 在metric处理完成后提供一个signal机制
type TrackingAccumulator interface {
	Accumulator

	// Add the Metric and arrange for tracking feedback after processing..
	AddTrackingMetric(m Metric) TrackingID

	// Add a group of Metrics and arrange for a signal when the group has been
	// processed.
	AddTrackingMetricGroup(group []Metric) TrackingID

	// Delivered returns a channel that will contain the tracking results.
	Delivered() <-chan DeliveryInfo
}
