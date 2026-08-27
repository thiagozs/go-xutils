package calc

import (
	"errors"
	"math"
	"testing"
)

func TestCalc(t *testing.T) {
	t.Run("Calculate Limit and Offset with Strings", func(t *testing.T) {
		pageNumberStr := "2"
		pageSizeStr := "10"
		expectedLimit := int32(10)
		expectedOffset := int32(10)

		limit, offset, err := LimitOffsetString(pageNumberStr, pageSizeStr)
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

		limit, offset, err := LimitOffset(pageNumber, pageSize)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if limit != expectedLimit || offset != expectedOffset {
			t.Errorf("expected limit: %d, offset: %d, but got limit: %d, offset: %d", expectedLimit, expectedOffset, limit, offset)
		}
	})

	t.Run("Error on Invalid String Inputs", func(t *testing.T) {
		pageNumberStr := "invalid"
		pageSizeStr := "10"

		_, _, err := LimitOffsetString(pageNumberStr, pageSizeStr)
		if err == nil {
			t.Errorf("expected error, but got nil")
		}
	})

	t.Run("Error on Negative Page Number", func(t *testing.T) {
		pageNumberStr := "-1"
		pageSizeStr := "10"

		_, _, err := LimitOffsetString(pageNumberStr, pageSizeStr)
		if err == nil {
			t.Errorf("expected error, but got nil")
		}
	})

	t.Run("Error on Negative Page Size", func(t *testing.T) {
		pageNumberStr := "1"
		pageSizeStr := "-10"

		_, _, err := LimitOffsetString(pageNumberStr, pageSizeStr)
		if err == nil {
			t.Errorf("expected error, but got nil")
		}
	})

	t.Run("Error on Zero Page Number", func(t *testing.T) {
		pageNumberStr := "0"
		pageSizeStr := "10"

		_, _, err := LimitOffsetString(pageNumberStr, pageSizeStr)
		if err == nil {
			t.Errorf("expected error, but got nil")
		}
	})

	t.Run("Error on Zero Page Size", func(t *testing.T) {
		pageNumberStr := "1"
		pageSizeStr := "0"

		_, _, err := LimitOffsetString(pageNumberStr, pageSizeStr)
		if err == nil {
			t.Errorf("expected error, but got nil")
		}
	})

	t.Run("Error on Invalid Int32 Inputs", func(t *testing.T) {
		pageNumber := int32(-1)
		pageSize := int32(10)

		_, _, err := LimitOffset(pageNumber, pageSize)
		if err == nil {
			t.Errorf("expected error, but got nil")
		}
	})

	t.Run("Error on Zero Int32 Inputs", func(t *testing.T) {
		pageNumber := int32(0)
		pageSize := int32(10)

		_, _, err := LimitOffset(pageNumber, pageSize)
		if err == nil {
			t.Errorf("expected error, but got nil")
		}
	})

	t.Run("Error on Negative Int32 Inputs", func(t *testing.T) {
		pageNumber := int32(-1)
		pageSize := int32(10)

		_, _, err := LimitOffset(pageNumber, pageSize)
		if err == nil {
			t.Errorf("expected error, but got nil")
		}
	})

	t.Run("Random In Range with Strings", func(t *testing.T) {
		minStr := "10"
		maxStr := "20"
		min := int32(10)
		max := int32(20)

		random, err := RandomInt32String(minStr, maxStr)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if random < min || random > max {
			t.Errorf("expected random number between %d and %d, but got %d", min, max, random)
		}
	})

	t.Run("Random In Range with Int32", func(t *testing.T) {
		min := int32(10)
		max := int32(20)

		random, err := RandomInt32(min, max)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if random < min || random > max {
			t.Errorf("expected random number between %d and %d, but got %d", min, max, random)
		}
	})

}

func TestBoundaries(t *testing.T) {
	got, err := RandomInt32(math.MinInt32, math.MaxInt32)
	if err != nil {
		t.Fatalf("full int32 range failed: value=%d err=%v", got, err)
	}

	_, _, err = LimitOffset(math.MaxInt32, 2)
	if !errors.Is(err, ErrOverflow) {
		t.Fatalf("expected ErrOverflow, got %v", err)
	}

	if _, _, err := LimitOffsetString("2147483648", "1"); err == nil {
		t.Fatal("expected int32 overflow to be rejected")
	}
	if _, _, err := LimitOffsetString("-2147483649", "1"); err == nil {
		t.Fatal("expected int32 underflow to be rejected")
	}
}
