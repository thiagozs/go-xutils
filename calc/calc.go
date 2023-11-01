package calc

import (
	"fmt"
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
