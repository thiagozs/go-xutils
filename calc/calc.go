package calc

import (
	"fmt"
	"strconv"
)

type Calc struct{}

func New() *Calc {
	return &Calc{}
}

// CalculateLimitAndOffset calculates limit and offset
func (c Calc) CalculateLimitAndOffsetStr(pageNumberStr, pageSizeStr string) (limit, offset int32, err error) {
	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil {
		return 0, 0, fmt.Errorf("calc: error converting pageNumberStr to int: %v", err)
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		return 0, 0, fmt.Errorf("calc: error converting pageSizeStr to int: %v", err)
	}

	if pageNumber < 1 || pageSize < 1 {
		return 0, 0, fmt.Errorf("calc: pageNumber and pageSize must be greater than 0")
	}

	limit, offset = c.CalculateLimitAndOffsetInt32(int32(pageNumber), int32(pageSize))
	return limit, offset, nil
}

// CalculateLimitAndOffsetInt32 calculates limit and offset
func (c Calc) CalculateLimitAndOffsetInt32(pageNumber, pageSize int32) (limit, offset int32) {
	if pageNumber < 1 {
		pageNumber = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	limit = int32(pageSize)
	offset = int32((pageNumber - 1) * pageSize)

	return limit, offset
}
