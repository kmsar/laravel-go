package Slices

import "math"

// ArrayChunk splits an array into chunks, returns nil if size < 1
func ArrayChunk[T comparable](array []T, size int) [][]T {
	if size < 1 {
		return nil
	}
	length := len(array)
	chunkNum := int(math.Ceil(float64(length) / float64(size)))
	var chunks [][]T
	for i, end := 0, 0; chunkNum > 0; chunkNum-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		chunks = append(chunks, array[i*size:end])
		i++
	}
	return chunks
}
