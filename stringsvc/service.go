package stringsvc

import "context"

type StringService interface {
	UpperString(c context.Context, s string) (string, error)
	Count(s string) (int, error)
}
