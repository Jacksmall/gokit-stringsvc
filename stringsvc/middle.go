package stringsvc

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/log"
)

// LoggingMiddleware struct
type LoggingMiddleware struct {
	Logger log.Logger
	Next   StringService
}

func (mw LoggingMiddleware) UpperString(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.UpperString(s)
	return
}

func (mw LoggingMiddleware) Count(s string) (n int, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "count",
			"input", s,
			"output", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n, err = mw.Next.Count(s)
	return
}

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           StringService
}

func (imw InstrumentingMiddleware) UpperString(s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "uppercase", "error", fmt.Sprint(err != nil)}
		imw.RequestCount.With(lvs...).Add(1)
		imw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = imw.Next.UpperString(s)
	return
}

func (imw InstrumentingMiddleware) Count(s string) (n int, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "count", "error", "false"}
		imw.RequestCount.With(lvs...).Add(1)
		imw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		imw.CountResult.Observe(float64(n))
	}(time.Now())

	n, err = imw.Next.Count(s)
	return
}
