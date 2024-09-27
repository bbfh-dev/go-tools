package tools

func SlicesDiff[T string | int | bool](a, b []T) []T {
	if len(a) < len(b) {
		return SlicesDiff(b, a)
	}

	m := make(map[T]bool)
	for _, item := range b {
		m[item] = true
	}

	var diff []T
	for _, item := range a {
		if _, found := m[item]; !found {
			diff = append(diff, item)
		}
		m[item] = false
	}

	for key, value := range m {
		if value {
			diff = append(diff, key)
		}
	}

	return diff
}
