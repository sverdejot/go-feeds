package validate

import (
	"strings"
	"errors"
	"fmt"
)

var ErrInvalidPrefix = func(Prefix string) error {
	return errors.New(fmt.Sprintf("the provided filename does not starts with %s", Prefix))
}

var ErrInvalidSuffix = func(Suffix string) error {
	return errors.New(fmt.Sprintf("the provided filename does not starts with %s", Suffix))
}

type EndsWith struct {
	Suffix		string
}

func (e EndsWith) Eval(filename string) error {
	if !strings.HasSuffix(filename, e.Suffix) {
		return errors.New("invalid")
	}
	return nil
}

type StartsWith struct {
	Prefix	string
}

func (s StartsWith) Eval(filename string) error {
	if !strings.HasPrefix(filename, s.Prefix) {
		return ErrInvalidPrefix(s.Prefix)
	}
	return nil
}
