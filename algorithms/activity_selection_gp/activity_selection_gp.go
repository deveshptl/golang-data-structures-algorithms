package main

import "fmt"

type activity struct {
	start  float64
	finish float64
}

// ActivitySelection selects the list of activities from given array
func ActivitySelection() {
	j := 0
	for i := range activities {
		if i == 0 {
			selected = append(selected, i)
			j = i
		} else if activities[j].finish < activities[i].start {
			selected = append(selected, i)
			j = i
		}
	}
}

var n int
var activities []activity
var selected []int

func init() {
	selected = make([]int, 0)
}

func main() {
	fmt.Println("\n-- Activity Selection Problem using Greedy Approach --")

	fmt.Print("\nEnter the number of activities: ")
	fmt.Scanf("%d", &n)
	activities = make([]activity, n)

	fmt.Println("Start entering starting and finishing times of each activities:")
	var temp float64
	for i := 0; i < n; i++ {
		fmt.Print("Start: ")
		fmt.Scanf("%f", &temp)
		activities[i].start = temp
		fmt.Print("Finish: ")
		fmt.Scanf("%f", &temp)
		activities[i].finish = temp
	}

	quicksort(0, n-1)

	ActivitySelection()

	fmt.Println("Selected activities are:")
	for i := range selected {
		fmt.Println(activities[selected[i]].start, activities[selected[i]].finish)
	}
}

// activities = []activity{
// activity{start: 1, finish: 4},
// activity{start: 3, finish: 5},
// activity{start: 0, finish: 6},
// activity{start: 5, finish: 7},
// activity{start: 3, finish: 8},
// activity{start: 5, finish: 9},
// activity{start: 6, finish: 10},
// activity{start: 8, finish: 11},
// activity{start: 8, finish: 12},
// activity{start: 2, finish: 13},
// activity{start: 12, finish: 14},
// }
// n = 11
