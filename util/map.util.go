package util

func GetMapKeys[T comparable, K any](list map[T]K) []T {
	ans := []T{}
	for key := range list {
		ans = append(ans, key)
	}
	return ans
}
