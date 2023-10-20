package calc

import (
	"testing"
)

func TestCalc(t *testing.T) {
	c := New()

	t.Run("Calculate Limit and Offset with Strings", func(t *testing.T) {
		pageNumberStr := "2"
		pageSizeStr := "10"
		expectedLimit := int32(10)
		expectedOffset := int32(10)

		limit, offset, err := c.CalculateLimitAndOffsetStr(pageNumberStr, pageSizeStr)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if limit != expectedLimit || offset != expectedOffset {
			t.Errorf("expected limit: %d, offset: %d, but got limit: %d, offset: %d", expectedLimit, expectedOffset, limit, offset)
		}
	})

	t.Run("Calculate Limit and Offset with Int32", func(t *testing.T) {
		pageNumber := int32(3)
		pageSize := int32(20)
		expectedLimit := int32(20)
		expectedOffset := int32(40)

		limit, offset := c.CalculateLimitAndOffsetInt32(pageNumber, pageSize)
		if limit != expectedLimit || offset != expectedOffset {
			t.Errorf("expected limit: %d, offset: %d, but got limit: %d, offset: %d", expectedLimit, expectedOffset, limit, offset)
		}
	})

	t.Run("Error on Invalid String Inputs", func(t *testing.T) {
		pageNumberStr := "invalid"
		pageSizeStr := "10"

		_, _, err := c.CalculateLimitAndOffsetStr(pageNumberStr, pageSizeStr)
		if err == nil {
			t.Errorf("expected error, but got nil")
		}
	})

	t.Run("Error on Negative Page Number", func(t *testing.T) {
		pageNumberStr := "-1"
		pageSizeStr := "10"

		_, _, err := c.CalculateLimitAndOffsetStr(pageNumberStr, pageSizeStr)
		if err == nil {
			t.Errorf("expected error, but got nil")
		}
	})
}
