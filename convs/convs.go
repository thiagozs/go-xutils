package convs

import (
	"fmt"
	"strconv"
)

// Parse converts a string to the requested primitive type.
func Parse[T any](s string) (T, error) {
	var zero T
	switch any(zero).(type) {
	case int:
		val, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			return zero, err
		}
		return any(int(val)).(T), nil
	case int8:
		val, err := strconv.ParseInt(s, 10, 8)
		if err != nil {
			return zero, err
		}
		return any(int8(val)).(T), nil
	case int16:
		val, err := strconv.ParseInt(s, 10, 16)
		if err != nil {
			return zero, err
		}
		return any(int16(val)).(T), nil
	case int32:
		val, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return zero, err
		}
		return any(int32(val)).(T), nil
	case int64:
		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return zero, err
		}
		return any(val).(T), nil
	case uint:
		val, err := strconv.ParseUint(s, 10, 0)
		if err != nil {
			return zero, err
		}
		return any(uint(val)).(T), nil
	case uint8:
		val, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			return zero, err
		}
		return any(uint8(val)).(T), nil
	case uint16:
		val, err := strconv.ParseUint(s, 10, 16)
		if err != nil {
			return zero, err
		}
		return any(uint16(val)).(T), nil
	case uint32:
		val, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			return zero, err
		}
		return any(uint32(val)).(T), nil
	case uint64:
		val, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return zero, err
		}
		return any(val).(T), nil
	case float32:
		val, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return zero, err
		}
		return any(float32(val)).(T), nil
	case float64:
		val, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return zero, err
		}
		return any(val).(T), nil
	case bool:
		val, err := strconv.ParseBool(s)
		if err != nil {
			return zero, err
		}
		return any(val).(T), nil
	case string:
		return any(s).(T), nil
	default:
		return zero, fmt.Errorf("unsupported type: %T", zero)
	}
}

// Format converts a primitive value to its string representation.
func Format[T any](input T) (string, error) {
	switch v := any(input).(type) {
	case int:
		return strconv.Itoa(v), nil
	case int8:
		return strconv.FormatInt(int64(v), 10), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int32:
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case uint:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint64:
		return strconv.FormatUint(v, 10), nil
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case bool:
		return strconv.FormatBool(v), nil
	default:
		return "", fmt.Errorf("unsupported type: %T", input)
	}
}
