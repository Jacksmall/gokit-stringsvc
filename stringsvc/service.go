package stringsvc

import (
	"errors"
	"strings"
)

type StringService interface {
	UpperCase(s string) (string, error)
	Count(s string) (int, error)
}

type stringService struct{}

func NewStringService() StringService {
	return &stringService{}
}

var ErrEmptyString = errors.New("empty string")

func (ss stringService) UpperCase(s string) (string, error) {
	if s == "" {
		return "", ErrEmptyString
	}
	return strings.ToUpper(s), nil
}

func (ss stringService) Count(s string) (int, error) {
	return len(s), nil
}
