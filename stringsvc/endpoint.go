package stringsvc

import (
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
