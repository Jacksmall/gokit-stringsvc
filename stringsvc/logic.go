package stringsvc

import (
	"context"
	"errors"
	"strings"
)

type stringService struct{}

var ErrEmptyString = errors.New("empty string")

func NewStringService() StringService {
	return &stringService{}
}

func (ss stringService) UpperString(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmptyString
	}
	return strings.ToUpper(s), nil
}

func (ss stringService) Count(s string) (int, error) {
	return len(s), nil
}
