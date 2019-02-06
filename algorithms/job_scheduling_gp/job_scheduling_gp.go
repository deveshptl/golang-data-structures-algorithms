package main

import "fmt"

type job struct {
	profit   float64
	deadline float64
}

// ScheduleJobs will schedule all the jobs maximizing the profit
func ScheduleJobs() {
	count := max
	i := 0
	deadline := -1
	for count != 0 {
		if i == len(jobs)-1 { // a complete scan of all jobs is completed
			break
		}
		if deadline == -1 { // if deadline gets to -1 it can't be scheduled
			deadline = int(jobs[i].deadline)
			if deadline == 0 && result[deadline-1] == 0 { // if a job's deadline is 0
				result[deadline] = i + 1
				deadline = -1
				count--
				i++
				continue
			}
		}
		if deadline-1 != -1 && result[deadline-1] == 0 { // check if the job can scheduled
			result[deadline-1] = i + 1
			deadline = -1
			count--
			i++
		} else if deadline-1 == -1 { // check if the deadline is 0
			deadline = -1
			i++
		} else if result[deadline-1] != 0 {
			deadline--
		}
	}
}

var n int
var jobs []job
var result []int
var max float64

func main() {
	fmt.Println("\n-- Job Scheduling using Greedy Programming --")

	fmt.Print("\nEnter the number of jobs: ")
	fmt.Scanf("%d", &n)
	jobs = make([]job, n)

	fmt.Println("Start entering profits and deadlines of each job:")
	var temp float64
	for i := 0; i < n; i++ {
		fmt.Print("Profit: ")
		fmt.Scanf("%f", &temp)
		jobs[i].profit = temp
		fmt.Print("Deadline: ")
		fmt.Scanf("%f", &temp)
		jobs[i].deadline = temp
		if i == 0 {
			max = jobs[i].deadline
		} else if max < jobs[i].deadline {
			max = jobs[i].deadline
		}
	}
	result = make([]int, int(max))
	quicksort(0, n-1)
	ScheduleJobs()
	total := float64(0)
	for i := range result {
		total += jobs[result[i]-1].profit
	}
	fmt.Println("Jobs", result, "will be scheduled to have", total, "units of profit and will be taken in ascending order of their deadlines.")
}
