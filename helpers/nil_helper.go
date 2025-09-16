package helpers

import "reflect"

func IsNilOrEmpty(x any) bool {
	if x == nil {
		return true
	}
	v := reflect.ValueOf(x)

	switch v.Kind() {
	case reflect.Interface, reflect.Pointer:
		if v.IsNil() {
			return true
		}
		// Optionally treat a pointer to a zero/empty value as empty
		return IsNilOrEmpty(v.Elem().Interface())

	case reflect.String, reflect.Array:
		return v.Len() == 0

	case reflect.Slice, reflect.Map:
		return v.IsNil() || v.Len() == 0

	default:
		// numbers, bools, structs, etc.
		return v.IsZero()
	}
}