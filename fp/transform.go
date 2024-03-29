package fp

// ReferenceValue creates a pointer to the given value.
//
// This is useful to create a reference to a value that is not a variable,
// such as an integer literal or the result of a function.
func ReferenceValue[T any](v T) *T {
	return &v
}

// DereferencePointer, unsurprisingly, dereferences a pointer.
//
// If the pointer is nil, the zero value of T is returned.
func DereferencePointer[T any](ptr *T) T {
	var value T

	if ptr == nil {
		return value
	}

	value = *ptr

	return value
}
