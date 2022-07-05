package main

import (
	"net/http"
	"os"

	"github.com/Jacksmall/gokit-stringsvc/stringsvc"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// logger
	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	// service layer
	svc := stringsvc.NewStringService()
	// loggingMiddleware struct
	svc = stringsvc.LoggingMiddleware{logger, svc}
	svc = stringsvc.InstrumentingMiddleware{requestCount, requestLatency, countResult, svc}
	// endpoints layer
	endpoints := stringsvc.MakeEndpoints(svc)

	// log middleware func
	// endpoints.UpperString = loggingMiddleware(log.With(logger, "method", "upperstring"))(endpoints.UpperString)
	// endpoints.Count = loggingMiddleware(log.With(logger, "method", "count"))(endpoints.Count)

	// transport layer
	upperStringHandler := httptransport.NewServer(
		endpoints.UpperString,
		stringsvc.DecodeUpperString,
		stringsvc.EncodeResponse,
	)
	countHandler := httptransport.NewServer(
		endpoints.Count,
		stringsvc.DecodeCountRequest,
		stringsvc.EncodeResponse,
	)
	// http
	http.Handle("/uppercase", upperStringHandler)
	http.Handle("/count", countHandler)
	http.Handle("/metrics", promhttp.Handler())

	logger.Log("msg", "HTTP", "addr", ":8082")
	logger.Log("err", http.ListenAndServe(":8082", nil))
}

// ============== loggingMiddleware move to the middleware =============
// func loggingMiddleware(logger log.Logger) endpoint.Middleware {
// 	return func(next endpoint.Endpoint) endpoint.Endpoint {
// 		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
// 			logger.Log("msg", "calling endpoint")
// 			defer logger.Log("msg", "calling endpoint")
// 			return next(ctx, request)
// 		}
// 	}
// }
// ============== loggingMiddleware move to the middleware =============

// ============== decode&&encode move to transport ==============
// func decodeUpperString(_ context.Context, r *http.Request) (interface{}, error) {
// 	var req stringsvc.UpperStringRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		return nil, err
// 	}
// 	return req, nil
// }

// func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
// 	var req stringsvc.CountRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		return nil, err
// 	}
// 	return req, nil
// }

// func encodeResponse(_ context.Context, w http.ResponseWriter, r interface{}) error {
// 	return json.NewEncoder(w).Encode(&r)
// }
// ============== decode&&encode move to transport ==============
