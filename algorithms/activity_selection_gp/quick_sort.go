package main

func partition(first, last int) int {
	pivot := activities[first].finish
	i := first
	j := last + 1

	for i < j {
		i++
		j--
		for activities[i].finish < pivot && i < last {
			i++
		}
		for activities[j].finish > pivot && j > first {
			j--
		}
		if i < j {
			activities[i].finish, activities[j].finish = activities[j].finish, activities[i].finish
		}
	}
	activities[j].finish, activities[first].finish = activities[first].finish, activities[j].finish
	return j
}

func quicksort(first, last int) {
	var sp int
	if first < last {
		sp = partition(first, last)
		quicksort(first, sp-1)
		quicksort(sp+1, last)
	}
}
