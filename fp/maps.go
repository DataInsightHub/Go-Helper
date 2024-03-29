package fp

// Values returns all values of the given map as a slice in a random order.
func Values[T any, K comparable](m map[K]T) []T {
	values := make([]T, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

// Keys returns all keys of the given map as a slice in a random order.
func Keys[T any, K comparable](m map[K]T) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

// CreateResultMap creates a new map, where all keys have been transformed by the given function.
func CreateMap[K comparable, T any](keyFunc func(T) K, values []T) map[K]T {
	m := make(map[K]T, len(values))
	for _, value := range values {
		m[keyFunc(value)] = value
	}
	return m
}

// MapMap creates a new map, where all values have been transformed by the given function.
func MapMap[K comparable, T any, V any](m map[K]T, fn func(T) V) map[K]V {
	result := make(map[K]V)
	for key, value := range m {
		result[key] = fn(value)
	}
	return result
}

// MapKey creates a new map, where all keys have been transformed by the given function.
//
// If fn maps two old keys to the same new key, one value is overwritten.
func MapKey[K1, K2 comparable, V any](m map[K1]V, fn func(K1) K2) map[K2]V {
	result := make(map[K2]V)
	for key, value := range m {
		result[fn(key)] = value
	}
	return result
}
