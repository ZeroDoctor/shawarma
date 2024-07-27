package service

type Comparable interface {
	GetID() int
}

func NewSet[T Comparable](values ...T) map[int]T {
	set := make(map[int]T, len(values))

	for i := range values {
		set[values[i].GetID()] = values[i]
	}

	return set
}

func AppendSet[T Comparable](source map[int]T, values ...T) map[int]T {
	for i := range values {
		source[values[i].GetID()] = values[i]
	}

	return source
}

func SetToSlice[T Comparable](source map[int]T) []T {
	var slice []T

	for _, value := range source {
		slice = append(slice, value)
	}

	return slice
}
