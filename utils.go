package errtmpl

import (
	"reflect"
)

func isEmpty(v interface{}) bool {
	switch t := v.(type) {
	case int, int8, int16, int32, int64:
		if t == 0 {
			return true
		}
		break
	case uint, uint8, uint16, uint32, uint64:
		if t == 0 {
			return true
		}
		break
	case float32, float64:
		if t == 0.0 {
			return true
		}
		break
	case string:
		if t == "" {
			return true
		}
		break
		//DO NOT SUPPORT NOW
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(v)
		if s.Len() == 0 {
			return true
		}
		break
	}

	return false
}
