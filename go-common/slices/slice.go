package slices

// SliceToMap takes a slice of structs and returns a map where the key is derived from a field in the struct.
func SliceToMap[K comparable, V any](slice []V, keyFunc func(V) K) map[K][]V {
	result := make(map[K][]V)
	for _, item := range slice {
		key := keyFunc(item)
		curr := result[key]
		curr = append(curr, item)
		result[key] = curr
	}
	return result
}
