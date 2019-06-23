package main

func compareX(p1, p2 Point) bool {
	return (p1.x < p2.x) || (p1.x == p2.x && p1.y > p2.y)
}

func compareY(p1, p2 Point) bool {
	return (p1.y < p2.y) || (p1.y == p2.y && p1.x < p2.x)
}

// Partition splits the arr into two arrays from pivot element
func Partition(pts []Point, first, last int, compare func(p1, p2 Point) bool) int {
	pivot := pts[first]
	i := first
	j := last + 1

	for i < j {
		i++
		j--
		for compare(pts[i], pivot) && i < last {
			i++
		}
		for !compare(pts[j], pivot) && j > first {
			j--
		}
		if i < j {
			pts[i], pts[j] = pts[j], pts[i]
		}
	}
	pts[j], pts[first] = pts[first], pts[j]
	return j
}

// QuickSort algorithms sorts the array using the compare function passed
func QuickSort(pts []Point, first, last int, compare func(p1, p2 Point) bool) {
	var sp int
	if first < last {
		sp = Partition(pts, first, last, compare)
		QuickSort(pts, first, sp-1, compare)
		QuickSort(pts, sp+1, last, compare)
	}
}
