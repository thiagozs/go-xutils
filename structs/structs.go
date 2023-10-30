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
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Use JSON tag as field name if it exists, otherwise use struct field name
		name := fieldType.Tag.Get("json")
		if name == "" {
			name = fieldType.Name
		}

		// Convert field name to lower case to follow common query parameter naming conventions
		name = strings.ToLower(name)

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
	return query.Encode()
}
