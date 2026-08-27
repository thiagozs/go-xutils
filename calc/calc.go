package calc

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/thiagozs/go-xutils/v2/randutil"
)

var (
	ErrInvalidRange = errors.New("calc: min must be less than or equal to max")
	ErrInvalidPage  = errors.New("calc: page number and page size must be greater than zero")
	ErrOverflow     = errors.New("calc: result overflows int32")
)

// RandomInt32 returns a non-cryptographic random value in the inclusive range.
func RandomInt32(min, max int32) (int32, error) {
	return randomInRange(int32Value(min), int32Value(max))
}

// RandomInt32String parses the bounds and returns an inclusive random value.
func RandomInt32String(min, max string) (int32, error) {
	return randomInRange(stringValue(min), stringValue(max))
}

func randomInRange[T convertible](min, max T) (int32, error) {
	minInt, err := min.ToInt32()
	if err != nil {
		return 0, fmt.Errorf("calc: error converting min to int32: %w", err)
	}

	maxInt, err := max.ToInt32()
	if err != nil {
		return 0, fmt.Errorf("calc: error converting max to int32: %w", err)
	}

	if minInt > maxInt {
		return 0, ErrInvalidRange
	}

	// int64 prevents overflow for the complete int32 range.
	width := int64(maxInt) - int64(minInt) + 1
	return int32(int64(minInt) + randutil.Default().Int63n(width)), nil
}

// LimitOffset calculates SQL-style pagination values.
func LimitOffset(pageNumber, pageSize int32) (int32, int32, error) {
	return limitOffset(int32Value(pageNumber), int32Value(pageSize))
}

// LimitOffsetString parses pagination values before calculating them.
func LimitOffsetString(pageNumber, pageSize string) (int32, int32, error) {
	return limitOffset(stringValue(pageNumber), stringValue(pageSize))
}

func limitOffset[T convertible](pageNumber, pageSize T) (int32, int32, error) {
	pageNumberInt, err := pageNumber.ToInt32()
	if err != nil {
		return 0, 0, fmt.Errorf("calc: error converting pageNumber to int32: %w", err)
	}

	pageSizeInt, err := pageSize.ToInt32()
	if err != nil {
		return 0, 0, fmt.Errorf("calc: error converting pageSize to int32: %w", err)
	}

	if pageNumberInt < 1 || pageSizeInt < 1 {
		return 0, 0, ErrInvalidPage
	}

	limit := pageSizeInt
	offset64 := int64(pageNumberInt-1) * int64(pageSizeInt)
	if offset64 > math.MaxInt32 {
		return 0, 0, ErrOverflow
	}
	offset := int32(offset64)

	return limit, offset, nil
}

type convertible interface {
	ToInt32() (int32, error)
}

type stringValue string

func (s stringValue) ToInt32() (int32, error) {
	i, err := strconv.ParseInt(string(s), 10, 32)
	return int32(i), err
}

type int32Value int32

func (i int32Value) ToInt32() (int32, error) {
	return int32(i), nil
}
