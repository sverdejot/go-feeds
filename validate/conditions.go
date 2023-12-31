package validate

import (
	"strings"
	"errors"
	"fmt"
)

var ErrInvalidPrefix = func(prefix string) error {
	return errors.New(fmt.Sprintf("the provided filename does not starts with %s", prefix))
}

var ErrInvalidSuffix = func(suffix string) error {
	return errors.New(fmt.Sprintf("the provided filename does not starts with %s", suffix))
}

type EndsWith struct {
	suffix		string
}

func (e EndsWith) Eval(filename string) error {
	if !strings.HasSuffix(filename, e.suffix) {
		return errors.New("invalid")
	}
	return nil
}

type StartsWith struct {
	prefix	string
}

func (s StartsWith) Eval(filename string) error {
	if !strings.HasPrefix(filename, s.prefix) {
		return ErrInvalidPrefix(s.prefix)
	}
	return nil
}
