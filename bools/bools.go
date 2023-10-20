package bools

import (
	"fmt"
	"strings"
)

type Bools struct{}

func New() *Bools {
	return &Bools{}
}

func (b *Bools) ToBool(s string) (bool, error) {
	switch strings.ToLower(s) {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, fmt.Errorf("invalid input: %s", s)
	}
}

func (b *Bools) ToString(bo bool) string {
	if bo {
		return "true"
	}
	return "false"
}
