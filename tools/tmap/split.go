package tmap

func Split[K comparable, V any](in map[K]V) ([]K, []V) {
	var outKeys []K
	var outValues []V

	for key, value := range in {
		outKeys = append(outKeys, key)
		outValues = append(outValues, value)
	}

	return outKeys, outValues
}

func Flatten[K comparable, V any](in map[K]V, format func(k K, v V) string) []string {
	var out = make([]string, len(in))
	var i = 0

	for key, value := range in {
		out[i] = format(key, value)
		i++
	}

	return out
}
