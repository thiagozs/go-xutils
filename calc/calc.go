package calc

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Calc struct{}

func New() *Calc {
	return &Calc{}
}

func (c Calc) CalculateLimitAndOffset(pageNumber, pageSize int32) (int32, int32, error) {
	return CalculateLimitAndOffset(Int32Convertible(pageNumber), Int32Convertible(pageSize))
}

func (c Calc) CalculateLimitAndOffsetStr(pageNumber, pageSize string) (int32, int32, error) {
	return CalculateLimitAndOffset(StringConvertible(pageNumber), StringConvertible(pageSize))
}

func (c Calc) RandomInRange(min, max int32) (int32, error) {
	return RandomInRange(Int32Convertible(min), Int32Convertible(max))
}

func (c Calc) RandomInRangeStr(min, max string) (int32, error) {
	return RandomInRange(StringConvertible(min), StringConvertible(max))
}

func RandomInRange[T Convertible](min, max T) (int32, error) {
	minInt, err := min.ToInt32()
	if err != nil {
		return 0, fmt.Errorf("calc: error converting min to int32: %w", err)
	}

	maxInt, err := max.ToInt32()
	if err != nil {
		return 0, fmt.Errorf("calc: error converting max to int32: %w", err)
	}

	if minInt > maxInt {
		return 0, fmt.Errorf("calc: min must be less than max")
	}

	return minInt + rand.Int31n(maxInt-minInt+1), nil
}

func CalculateLimitAndOffset[T Convertible](pageNumber, pageSize T) (int32, int32, error) {
	pageNumberInt, err := pageNumber.ToInt32()
	if err != nil {
		return 0, 0, fmt.Errorf("calc: error converting pageNumber to int32: %w", err)
	}

	pageSizeInt, err := pageSize.ToInt32()
	if err != nil {
		return 0, 0, fmt.Errorf("calc: error converting pageSize to int32: %w", err)
	}

	if pageNumberInt < 1 || pageSizeInt < 1 {
		return 0, 0, fmt.Errorf("calc: pageNumber and pageSize must be greater than 0")
	}

	limit := pageSizeInt
	offset := (pageNumberInt - 1) * pageSizeInt

	return limit, offset, nil
}

type Convertible interface {
	ToInt32() (int32, error)
}

type StringConvertible string

func (s StringConvertible) ToInt32() (int32, error) {
	i, err := strconv.Atoi(string(s))
	return int32(i), err
}

type Int32Convertible int32

func (i Int32Convertible) ToInt32() (int32, error) {
	return int32(i), nil
}
