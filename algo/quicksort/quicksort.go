package quicksort

func Do(l []int) []int {
	return do(l, 0, len(l)-1)
}

func do(l []int, left int, right int) []int {
	if left >= right {
		return l
	}

	p := (left + right) / 2
	index, l := partition(l, left, right, l[p])

	l = do(l, 0, index)
	l = do(l, index+1, right)

	return l
}

func partition(l []int, left int, right int, pivot int) (int, []int) {
	for left < right {
		for l[left] < pivot {
			left++
		}

		for l[right] > pivot {
			right--
		}

		if l[left] > l[right] {
			l[left], l[right] = l[right], l[left]
		}
	}

	return right, l
}
