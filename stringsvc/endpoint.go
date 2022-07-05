package stringsvc

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	UpperString endpoint.Endpoint
	Count       endpoint.Endpoint
}

func MakeEndpoints(s StringService) *Endpoints {
	return &Endpoints{
		UpperString: makeUpperString(s),
		Count:       makeCount(s),
	}
}

func makeUpperString(s StringService) endpoint.Endpoint {
	return func(c context.Context, request interface{}) (interface{}, error) {
		req := request.(UpperStringRequest)
		resp, err := s.UpperString(c, req.S)
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
