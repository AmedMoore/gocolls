package slice

// ForEachWithStop applies a provided function to each element in a slice of any type.
// It allows the function to stop iteration by invoking the provided stop function.
func ForEachWithStop[T any](slice []T, iterator func(item T, stop func())) {
	if iterator == nil {
		return
	}
	stopped := false
	stop := func() { stopped = true }
	for _, item := range slice {
		iterator(item, stop)
		if stopped {
			break
		}
	}
}
