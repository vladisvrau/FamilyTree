package util

func InterfacedList[T any](list []T) []any {
	result := []any{}

	for _, obj := range list {
		result = append(result, obj)
	}

	return result
}

func SlicePopComparable[T comparable](list []T, v T) []T {
	index := IndexOf(list, v)
	if index < 0 {
		return list
	}
	return append(list[:index], list[index+1:]...)
}

func IndexOf[T comparable](list []T, v T) int {
	for index, item := range list {
		if item == v {
			return index
		}
	}

	return -1
}
