package structs

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

var ErrExpectedStruct = errors.New("structs: expected a struct or non-nil pointer to struct")

// EncodeQuery converts exported struct fields to URL query parameters. It
// honors json names, "-", and omitempty and returns errors for invalid input.
func EncodeQuery(input any) (string, error) {
	if input == nil {
		return "", ErrExpectedStruct
	}
	v := reflect.ValueOf(input)
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return "", ErrExpectedStruct
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return "", ErrExpectedStruct
	}

	query := url.Values{}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)
		if !fieldType.IsExported() {
			continue
		}
		tag := fieldType.Tag.Get("json")
		name := fieldType.Name
		options := ""
		hasJSONName := false

		if tag != "" {
			parts := strings.Split(tag, ",")
			name = parts[0]
			hasJSONName = name != ""
			if len(parts) > 1 {
				options = "," + strings.Join(parts[1:], ",") + ","
			}
		}
		if name == "-" {
			continue
		}
		if name == "" {
			name = fieldType.Name
		}

		// Skip zero values for fields with omitempty
		if strings.Contains(options, ",omitempty,") && isEmptyValue(field) {
			continue
		}
		if !hasJSONName {
			name = strings.ToLower(name)
		}

		switch field.Kind() {
		case reflect.Slice:
			var sliceValues []string
			for j := 0; j < field.Len(); j++ {
				sliceValues = append(sliceValues, fmt.Sprintf("%v", field.Index(j)))
			}
			query.Add(name, strings.Join(sliceValues, ","))
		default:
			query.Add(name, fmt.Sprintf("%v", field.Interface()))
		}
	}
	return query.Encode(), nil
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
	case reflect.Interface, reflect.Pointer:
		return v.IsNil()
	}
	return false
}
