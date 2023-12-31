package convs

import (
	"fmt"
	"strconv"
)

type Convs struct{}

func New() *Convs {
	return &Convs{}
}

func (c *Convs) ToInt(input string) (int, error) {
	return NewConverter[int]().StringToType(input)
}

func (c *Convs) ToInt32(input string) (int32, error) {
	return NewConverter[int32]().StringToType(input)
}

func (c *Convs) ToInt64(input string) (int64, error) {
	return NewConverter[int64]().StringToType(input)
}

func (c *Convs) ToFloat32(input string) (float32, error) {
	return NewConverter[float32]().StringToType(input)
}

func (c *Convs) ToFloat64(input string) (float64, error) {
	return NewConverter[float64]().StringToType(input)
}

func (c *Convs) ToBool(input string) (bool, error) {
	return NewConverter[bool]().StringToType(input)
}

func (c *Convs) ToString(input any) (string, error) {
	return NewConverter[any]().ToString(input)
}

type Converter[T any] struct{}

func NewConverter[T any]() *Converter[T] {
	return &Converter[T]{}
}

func (c *Converter[T]) StringToType(s string) (T, error) {
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

func (c *Converter[T]) ToString(input T) (string, error) {
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
