package stringsvc

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type (
	UpperStringRequest struct {
		S string `json:"s"`
	}
	UpperStringResponse struct {
		V   string `json:"v"`
		Err string `json:"err"`
	}
	CountRequest struct {
		S string `json:"s"`
	}
	CountResponse struct {
		V int `json:"v"`
	}
)

func makeUpperString(s StringService) endpoint.Endpoint {
	return func(c context.Context, request interface{}) (interface{}, error) {
		req := request.(UpperStringRequest)
		resp, err := s.UpperString(req.S)
		if err != nil {
			return UpperStringResponse{resp, err.Error()}, err
		}
		return UpperStringResponse{resp, ""}, nil
	}
}

func makeCount(s StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		v, err := s.Count(req.S)
		if err != nil {
			return CountResponse{v}, err
		}
		return CountResponse{v}, nil
	}
}

func DecodeUpperString(_ context.Context, r *http.Request) (interface{}, error) {
	var req UpperStringRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, r interface{}) error {
	return json.NewEncoder(w).Encode(&r)
}
