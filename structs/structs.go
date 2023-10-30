package structs

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type Structs struct{}

func New() *Structs {
	return &Structs{}
}

func (s *Structs) ToQueryParams(i any) string {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return ""
	}

	query := url.Values{}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)
		tag := fieldType.Tag.Get("json")
		name := fieldType.Name

		// Use JSON tag as field name if it's available
		if tag != "" {
			name = strings.Split(tag, ",")[0] // Ignore options like omitempty
		}

		// Skip zero values for fields with omitempty
		if strings.Contains(tag, "omitempty") && isEmptyValue(field) {
			continue
		}

		switch field.Kind() {
		case reflect.Slice:
			var sliceValues []string
			for j := 0; j < field.Len(); j++ {
				sliceValues = append(sliceValues, fmt.Sprintf("%v", field.Index(j)))
			}
			query.Add(strings.ToLower(name), strings.Join(sliceValues, ","))
		default:
			query.Add(strings.ToLower(name), fmt.Sprintf("%v", field.Interface()))
		}
	}
	return query.Encode()
}

// isEmptyValue checks if a reflect.Value is considered empty
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
