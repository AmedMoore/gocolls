package slice

// Chunk takes a slice and splits it into chunks of a specified size.
func Chunk[T any](slice []T, chunkSize int) [][]T {
	if chunkSize == 0 {
		return make([][]T, 0)
	}
	chunks := make([][]T, (len(slice)+chunkSize-1)/chunkSize)
	for i := range chunks {
		chunks[i] = BoundedSub(slice, i*chunkSize, (i+1)*chunkSize)
	}
	return chunks
}
