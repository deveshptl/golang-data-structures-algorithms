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

// Point represents a point on a 2d plane having x, y coords
type Point struct {
	x float64
	y float64
}

var arr []Point

func init() {

	arr = make([]Point, 0)

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
	fmt.Println("-- Convex Hull Brute Force Technique --")

	// sort the points based on their x values and ties sorted according to y value
	MergeSort(0, len(arr)-1)

	// create a final convex hull points array
	ch := make([]Point, 0)

	// farthest two points are obviously the convex points. add them to 'ch'
	ch = append(ch, arr[0], arr[len(arr)-1])

	// call convex hull excluding those two points
	ch = append(ch, ConvexHull(arr[1:len(arr)-1], ch[0], ch[1])...)

	ch = SortConvexPoints(ch)

	for i := 0; i < len(ch); i++ {
		fmt.Println(ch[i])
	}
}

// ConvexHull finds a convex polygon of given set of 2d points
func ConvexHull(chArr []Point, p1, p2 Point) []Point {

	ch := make([]Point, 0)
	sLeft := make([]Point, 0)
	sRight := make([]Point, 0)

	for i := 0; i < len(chArr); i++ {
		value := FindSide(p1, p2, chArr[i])
		if value < 0 {
			sLeft = append(sLeft, chArr[i])
		} else if value > 0 {
			sRight = append(sRight, chArr[i])
		}
	}

	// solve the upper hull
	sLeft, pLeft := SolveULHull(sLeft, p1, p2)
	// solve the lower hull
	sRight, pRight := SolveULHull(sRight, p1, p2)

	// pLeft, pRight are new convex points obtained by solving upper and lower hulls. add them to 'ch'
	ch = append(ch, pLeft, pRight)

	// stop if there are no more points to the left side
	if len(sLeft) > 0 {
		ch = append(ch, ConvexHull(sLeft, p1, pLeft)...)
	}

	// stop if there are no more points to the right side
	if len(sRight) > 0 {
		ch = append(ch, ConvexHull(sRight, p1, pRight)...)
	}

	return ch
}

// SolveULHull solves upper and lower hull
func SolveULHull(pts []Point, p1, p2 Point) ([]Point, Point) {

	var p Point
	dist := math.Inf(-1)

	// finding a farthest point from point p1
	for i := 0; i < len(pts); i++ {
		newDist := FindDist(p1, pts[i])
		if newDist > dist {
			p = pts[i]
			dist = newDist
		} else if newDist == dist {
			// TODO: Run a program for which this condition hold true just to make sure that it working as expected
			var angle1, angle2 float64
			// angle of previous
			angle1 = FindAngle(p, p1, p2)
			// angle of new
			angle2 = FindAngle(pts[i], p1, p2)
			if angle1 < angle2 {
				dist = newDist
				p = pts[i]
			}
		}
	}

	// removing all points which are in the triangle formed by p, p1, p2
	temp := make([]Point, 0)

	for i := 0; i < len(pts); i++ {
		if pts[i].x == p.x && pts[i].y == p.y {
			continue
		}
		if !IsInside(p1, p2, p, pts[i]) {
			temp = append(temp, pts[i])
		}
	}
	return temp, p
}

// Answer to points in sample.txt file
// (0.0, 0.0)
// (0.0, 3.0)
// (4.0, 0.0)
// (2.0, 9.0)
// (4.0, 4.0)
// (5.0, 1.0)

//
// Some other points
// 0 3
// 2 2
// 1 1
// 2 1
// 3 0
// 0 0
// 3 3
// -- Ans --
// (0, 3)
// (0, 0)
// (3, 0)
// (3, 3)
//
// Some other points
// 0 0
// 0 4
// -4 0
// 5 0
// 0 -6
// 1 0
// -- Ans --
// (-4, 0)
// (5, 0)
// (0, -6)
// (0, 4)
