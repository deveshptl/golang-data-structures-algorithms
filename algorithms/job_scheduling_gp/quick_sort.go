package main

// this quick sort algorithm will sort in descending order
func partition(first, last int) int {
	pivot := jobs[first].profit
	i := first
	j := last + 1

	for i < j {
		i++
		j--
		for jobs[i].profit > pivot && i < last {
			i++
		}
		for jobs[j].profit < pivot && j > first {
			j--
		}
		if i < j {
			jobs[i].profit, jobs[j].profit = jobs[j].profit, jobs[i].profit
		}
	}
	jobs[j].profit, jobs[first].profit = jobs[first].profit, jobs[j].profit
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
