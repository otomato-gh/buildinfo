package buildinfo

import (
	"reflect"
)

func checkRating(info BuildInfo) float32 {
	v := reflect.ValueOf(info)
	notEmptycount := 0
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface()
		if value != float32(0) && value != "" {
			notEmptycount++
		}
	}
	rating := 1.0 / float32(v.NumField()) * float32(notEmptycount)
	return float32(rating)
}
