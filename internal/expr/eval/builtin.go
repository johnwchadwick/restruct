package eval

import (
	"reflect"

	"github.com/go-restruct/restruct/internal/expr/value"
)

var (
	builtins = map[string]value.Value{
		"int":   value.NewFunc(BuiltinInt),
		"uint":  value.NewFunc(BuiltinUint),
		"float": value.NewFunc(BuiltinFloat),
		"len":   value.NewFunc(BuiltinLen),
		"first": value.NewFunc(BuiltinFirst),
		"last":  value.NewFunc(BuiltinLast),
		"sum":   value.NewFunc(BuiltinSum),
		"usum":  value.NewFunc(BuiltinUsum),
		"fsum":  value.NewFunc(BuiltinFsum),
	}
)

// BuiltinInt implements the builtin int function
func BuiltinInt(v interface{}) int64 {
	switch t := v.(type) {
	case int64:
		return int64(t)
	case uint64:
		return int64(t)
	case float64:
		return int64(t)
	default:
		panic("unexpected type")
	}
}

// BuiltinUint implements the builtin uint function
func BuiltinUint(v interface{}) uint64 {
	switch t := v.(type) {
	case int64:
		return uint64(t)
	case uint64:
		return uint64(t)
	case float64:
		return uint64(t)
	default:
		panic("unexpected type")
	}
}

// BuiltinFloat implements the builtin float function.
func BuiltinFloat(v interface{}) float64 {
	switch t := v.(type) {
	case int64:
		return float64(t)
	case uint64:
		return float64(t)
	case float64:
		return float64(t)
	default:
		panic("unexpected type")
	}
}

// BuiltinLen implements the builtin len function.
func BuiltinLen(v interface{}) uint64 {
	return uint64(reflect.ValueOf(v).Len())
}

// BuiltinFirst implements the naive first function.
func BuiltinFirst(v interface{}) interface{} {
	return reflect.ValueOf(v).Index(0).Interface()
}

// BuiltinLast implements the naive last function.
func BuiltinLast(v interface{}) interface{} {
	l := reflect.ValueOf(v).Len()
	return reflect.ValueOf(v).Index(l - 1).Interface()
}

// BuiltinSum implements the naive sum function.
func BuiltinSum(v ...interface{}) int64 {
	s := int64(0)
	for _, n := range v {
		switch t := n.(type) {
		case uint8:
			s += int64(t)
		case uint16:
			s += int64(t)
		case uint32:
			s += int64(t)
		case uint64:
			s += int64(t)
		case uint:
			s += int64(t)
		case int8:
			s += int64(t)
		case int16:
			s += int64(t)
		case int32:
			s += int64(t)
		case int64:
			s += int64(t)
		case int:
			s += int64(t)
		case float32:
			s += int64(t)
		case float64:
			s += int64(t)
		default:
			r := reflect.ValueOf(n)
			if r.Kind() == reflect.Ptr {
				s += BuiltinSum(r.Elem().Interface())
			} else if r.Kind() == reflect.Array || r.Kind() == reflect.Slice {
				for i := 0; i < r.Len(); i++ {
					s += BuiltinSum(r.Index(i).Elem().Interface())
				}
			} else if r.Kind() == reflect.Map {
				for _, key := range r.MapKeys() {
					s += BuiltinSum(r.MapIndex(key).Interface())
				}
			}
		}
	}
	return s
}

// BuiltinUsum implements the naive usum function.
func BuiltinUsum(v ...interface{}) uint64 {
	s := uint64(0)
	for _, n := range v {
		switch t := n.(type) {
		case uint8:
			s += uint64(t)
		case uint16:
			s += uint64(t)
		case uint32:
			s += uint64(t)
		case uint64:
			s += uint64(t)
		case uint:
			s += uint64(t)
		case int8:
			s += uint64(t)
		case int16:
			s += uint64(t)
		case int32:
			s += uint64(t)
		case int64:
			s += uint64(t)
		case int:
			s += uint64(t)
		case float32:
			s += uint64(t)
		case float64:
			s += uint64(t)
		default:
			r := reflect.ValueOf(n)
			if r.Kind() == reflect.Ptr {
				s += BuiltinUsum(r.Elem().Interface())
			} else if r.Kind() == reflect.Array || r.Kind() == reflect.Slice {
				for i := 0; i < r.Len(); i++ {
					s += BuiltinUsum(r.Index(i).Elem().Interface())
				}
			} else if r.Kind() == reflect.Map {
				for _, key := range r.MapKeys() {
					s += BuiltinUsum(r.MapIndex(key).Interface())
				}
			}
		}
	}
	return s
}

// BuiltinFsum implements the naive fsum function.
func BuiltinFsum(v ...interface{}) float64 {
	s := float64(0)
	for _, n := range v {
		switch t := n.(type) {
		case uint8:
			s += float64(t)
		case uint16:
			s += float64(t)
		case uint32:
			s += float64(t)
		case uint64:
			s += float64(t)
		case uint:
			s += float64(t)
		case int8:
			s += float64(t)
		case int16:
			s += float64(t)
		case int32:
			s += float64(t)
		case int64:
			s += float64(t)
		case int:
			s += float64(t)
		case float32:
			s += float64(t)
		case float64:
			s += float64(t)
		default:
			r := reflect.ValueOf(n)
			if r.Kind() == reflect.Ptr {
				s += BuiltinFsum(r.Elem().Interface())
			} else if r.Kind() == reflect.Array || r.Kind() == reflect.Slice {
				for i := 0; i < r.Len(); i++ {
					s += BuiltinFsum(r.Index(i).Elem().Interface())
				}
			} else if r.Kind() == reflect.Map {
				for _, key := range r.MapKeys() {
					s += BuiltinFsum(r.MapIndex(key).Interface())
				}
			}
		}
	}
	return s
}
