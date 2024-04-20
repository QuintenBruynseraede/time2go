package utils

func Repeat(value int, n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = value
	}
	return arr
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func Between(min int, max int) []int {
	arr := []int{}
	for i := min; i < max; i++ {
		arr = append(arr, i)
	}
	return arr
}
