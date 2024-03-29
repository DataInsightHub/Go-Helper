package fp

// The type Result represents the result of a operation that may fail.
//
// It mirrors the common pattern that a function returns a value and an error.
// Similarly, the Result type consists of a value of type T and an error (that may be nil).
//
// See [ResultFrom]
type Result[T any] interface {
	Ok() T
	Err() error
}

type result[T any] struct {
	ok  T
	err error
}

// Ok extracts the value from the result.
func (res result[T]) Ok() T {
	return res.ok
}

// Err extracts the error from the result.
//
// Returns nil if the operation was successful.
func (res result[T]) Err() error {
	return res.err
}

// ResultFrom creates a new Result.
func ResultFrom[T any](val T, err error) Result[T] {
	return result[T]{
		ok:  val,
		err: err,
	}
}
