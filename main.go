package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Jacksmall/gokit-stringsvc/stringsvc"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	// service layer
	svc := stringsvc.NewStringService()
	// endpoints layer
	endpoints := stringsvc.MakeEndpoints(svc)
	// transport layer
	upperStringHandler := httptransport.NewServer(
		endpoints.UpperString,
		decodeUpperString,
		encodeResponse,
	)
	countHandler := httptransport.NewServer(
		endpoints.Count,
		decodeCountRequest,
		encodeResponse,
	)
	// http
	http.Handle("/uppercase", upperStringHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func decodeUpperString(_ context.Context, r *http.Request) (interface{}, error) {
	var req stringsvc.UpperStringRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req stringsvc.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, r interface{}) error {
	return json.NewEncoder(w).Encode(&r)
}
