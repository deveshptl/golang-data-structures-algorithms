package main

func merge(l, m, r int) {
	var i, j, k int
	nl := m - l + 1
	nr := r - m
	var left, right []Point

	for i = 0; i < nl; i++ {
		left = append(left, arr[l+i])
	}

	for j = 0; j < nr; j++ {
		right = append(right, arr[m+1+j])
	}

	i = 0
	j = 0
	k = l

	for i < nl && j < nr {
		if (left[i].x < right[j].x) || (left[i].x == right[j].x && left[i].y < right[j].y) {
			arr[k] = left[i]
			i++
			k++
		} else {
			arr[k] = right[j]
			j++
			k++
		}
	}

	for i < nl {
		arr[k] = left[i]
		i++
		k++
	}
	for j < nr {
		arr[k] = right[j]
		j++
		k++
	}
}

// MergeSort sorts list of points in non decreasing order of x values and ties resolved with y values sorted in non decreasing order
func MergeSort(l, r int) {
	var mid int
	if l < r {
		mid = (l + r) / 2
		MergeSort(l, mid)
		MergeSort(mid+1, r)
		merge(l, mid, r)
	}
}
