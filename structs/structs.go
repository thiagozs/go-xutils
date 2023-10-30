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
		name := v.Type().Field(i).Name
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
