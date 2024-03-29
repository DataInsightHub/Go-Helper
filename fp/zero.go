package fp

import "reflect"

// IsZeroValue reports whether the given value is the zero value of its type.
func IsZeroValue[A any](v A) bool {

	value := reflect.ValueOf(v)

	if value.IsValid() {
		return value.IsZero()
	}

	return true
}

// ZeroValueOf returns the zero value of the given type.
func ZeroValueOf[A any]() A {

	var zeroValue A

	return zeroValue
}
