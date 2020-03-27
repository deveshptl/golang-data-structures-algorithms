package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Point represents a point on a 2d plane having x, y coords
type Point struct {
	x float64
	y float64
}

var arr []Point

// pointInConvex keeps a simple check
// whether a given point is considered in convex set or not
var pointInConvex map[string]bool

func init() {
	arr = make([]Point, 0)
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

		arr = append(arr, Point{x: strx, y: stry})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("\n-- Convex Hull Brute Force Technique --")
	ch := ConvexHull()
	ch = SortConvexPoints(ch)
	for i := 0; i < len(ch); i++ {
		fmt.Println(ch[i])
	}
}

// ConvexHull finds required set of points which forms a convex
func ConvexHull() []Point {
	newArr := make([]Point, 0)

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			flag := IsConvex(arr[i], arr[j])

			if !flag {
				continue
			}

			// float64 to string
			xi := fmt.Sprintf("%f\n", arr[i].x)
			yi := fmt.Sprintf("%f\n", arr[i].y)
			xj := fmt.Sprintf("%f\n", arr[j].x)
			yj := fmt.Sprintf("%f\n", arr[j].y)

			//string of x[i] y[i]
			strI := xi + yi
			//string of x[j] y[j]
			strJ := xj + yj

			if !pointInConvex[strI] {
				newArr = append(newArr, arr[i])
				pointInConvex[strI] = true
			}

			if !pointInConvex[strJ] {
				newArr = append(newArr, arr[j])
				pointInConvex[strJ] = true
			}
		}
	}
	return newArr
}

// IsConvex checks whether the given two points should be considered in the solution
func IsConvex(p1, p2 Point) bool {
	flag := float64(0)
	a := p2.y - p1.y
	b := p1.x - p2.x
	c := p1.x*p2.y - p1.y*p2.x

	for i := 0; i < len(arr); i++ {
		lhs := a*arr[i].x + b*arr[i].y - c
		if lhs == 0 {
			check := IsLongestPair(p1, p2, arr[i])
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
