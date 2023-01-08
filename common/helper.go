package common

func Difference[T comparable](list1 []T, list2 []T) (left []T, right []T) {
	leftMap := map[T]struct{}{}
	rightMap := map[T]struct{}{}

	for _, elem := range list1 {
		leftMap[elem] = struct{}{}
	}

	for _, elem := range list2 {
		rightMap[elem] = struct{}{}
	}

	for _, elem := range list1 {
		if _, ok := rightMap[elem]; !ok {
			left = append(left, elem)
		}
	}

	for _, elem := range list2 {
		if _, ok := leftMap[elem]; !ok {
			right = append(right, elem)
		}
	}

	return left, right
}

func Map[I, O any](in []I, f func(I) O) []O {
	out := make([]O, len(in))
	for i := range in {
		out[i] = f(in[i])
	}
	return out
}
