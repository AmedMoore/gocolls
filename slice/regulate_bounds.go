package slice

// RegulateBounds returns the regulated start and end bounds within the given length.
// If the start index is negative, it will be set to 0.
// If the end index is greater than length, it will be set to length.
// If the end index is less than the start index, it will be set equal to start.
func RegulateBounds(start int, end int, length int) (int, int) {
	if length < 1 {
		return 0, 0
	}
	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	if end < start {
		end = start
	}
	return start, end
}
