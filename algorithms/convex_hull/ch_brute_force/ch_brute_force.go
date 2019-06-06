package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var x []float64
var y []float64

// pointInConvex keeps a simple check
// whether a given point is considered in convex set or not
var pointInConvex map[string]bool

func init() {
	x = make([]float64, 0)
	y = make([]float64, 0)
	pointInConvex = make(map[string]bool)

	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		str := strings.Split(txt, " ")

		strx, err := strconv.ParseFloat(str[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		stry, err := strconv.ParseFloat(str[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		x = append(x, strx)
		y = append(y, stry)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("-- Convex Hull Brute Force Technique --")
	ax, ay := ConvexHull()
	for i := 0; i < len(ax); i++ {
		sx := fmt.Sprintf("%.1f", ax[i])
		sy := fmt.Sprintf("%.1f", ay[i])
		fmt.Println("(" + sx + ", " + sy + ")")
	}
}

// ConvexHull finds required set of points which forms a convex
func ConvexHull() ([]float64, []float64) {
	ax := make([]float64, 0)
	ay := make([]float64, 0)

	for i := 0; i < len(x)-1; i++ {
		for j := i + 1; j < len(x); j++ {
			flag := isConvex(x[i], y[i], x[j], y[j])

			// float64 to string
			xi := fmt.Sprintf("%f", x[i])
			yi := fmt.Sprintf("%f", y[i])
			xj := fmt.Sprintf("%f", x[j])
			yj := fmt.Sprintf("%f", y[j])

			//string of x[i] y[i]
			strI := xi + yi
			//string of x[j] y[j]
			strJ := xj + yj

			if flag && !pointInConvex[strI] {
				ax = append(ax, x[i])
				ay = append(ay, y[i])
				pointInConvex[strI] = true
			}

			if flag && !pointInConvex[strJ] {
				ax = append(ax, x[j])
				ay = append(ay, y[j])
				pointInConvex[strJ] = true
			}
		}
	}
	return ax, ay
}

// isConvex checks whether the given two points should be considered in the solution
func isConvex(x1, y1, x2, y2 float64) bool {
	flag := float64(0)
	a := y2 - y1
	b := x1 - x2
	c := x1*y2 - y1*x2

	for i := 0; i < len(x); i++ {
		lhs := a*x[i] + b*y[i] - c
		if lhs == 0 {
			check := isLongestPair([]float64{x1, y1}, []float64{x2, y2}, []float64{x[i], y[i]})
			if !check {
				return false
			}
		}
		if flag != 0 && lhs != 0 && !((lhs < 0 && flag < 0) || (lhs > 0 && flag > 0)) {
			return false
		}
		if flag == 0 {
			flag = lhs
		}
	}
	return true
}

// isLongestPair checks whether the given first two set of points (xy1, xy2) have longest distance or not
// The main idea is not to select all given points in question that lies on the same line
func isLongestPair(xy1, xy2, xy3 []float64) bool {
	d12 := math.Pow(xy1[0]-xy2[0], 2) + math.Pow(xy1[1]-xy2[1], 2)
	d13 := math.Pow(xy1[0]-xy3[0], 2) + math.Pow(xy1[1]-xy3[1], 2)
	d23 := math.Pow(xy2[0]-xy3[0], 2) + math.Pow(xy2[1]-xy3[1], 2)
	if d12 >= d13 && d12 >= d23 {
		return true
	}
	return false
}
