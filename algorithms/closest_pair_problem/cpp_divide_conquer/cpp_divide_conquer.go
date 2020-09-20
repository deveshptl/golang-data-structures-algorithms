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

// Note: All the functions that returns distance as one of the return values, is not actual the distance.
// The value is not squared root. It's just the sum of square of difference of x and y co-ordinates

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
	fmt.Println("\n-- Closest Pair Problem using Divide and Conquer Approach --")

	QuickSort(arr, 0, len(arr)-1, compareX)

	p, dist := ClosestPair(arr)
	fmt.Println("\nClosest pair of points are:", p[0], p[1], "whose distance is:", math.Sqrt(dist))
}

// ClosestPair finds two points in a given set of 2d points
// Returns a point and distance between those two points
func ClosestPair(pts []Point) ([]Point, float64) {

	if len(pts) <= 3 {
		return BruteForce(pts)
	}

	// minimum distance and points obtained from solving left side
	dLeft := math.Inf(1) // nolint:ineffassign
	pLeft := make([]Point, 2)

	// minimum distance and points obtained from solving right side
	dRight := math.Inf(1) // nolint:ineffassign
	pRight := make([]Point, 2)

	// minimum distance and points obtained from solving points near around the middle line
	distS := math.Inf(1) // nolint:ineffassign
	ps := make([]Point, 2)

	// final distance and points having the least distance
	dist := math.Inf(1) // nolint:ineffassign
	p := make([]Point, 2)

	// finding the median
	n := len(pts) / 2

	pLeft, dLeft = ClosestPair(pts[:n])
	pRight, dRight = ClosestPair(pts[n:])

	if dLeft < dRight {
		dist = dLeft
		p = pLeft
	} else if dRight < dLeft {
		dist = dRight
		p = pRight
	} else {
		fmt.Println("\nPoints with similar distance may exist.")
		dist = dLeft
		p = pLeft
	}

	// points near the middle of line having x coords difference with median point no more than minimum distance
	strip := make([]Point, 0)

	for i := 0; i < len(pts); i++ {
		if (math.Abs(pts[i].x) - pts[n].x) < dist {
			strip = append(strip, pts[i])
		}
	}

	// sort points of strip according to y coords1
	QuickSort(strip, 0, len(strip)-1, compareY)

	ps, distS = ClosestInStrip(strip, dist)

	if distS < dist {
		p = ps
		dist = distS
	}

	return p, dist
}

// FindDist finds distance between two points
func FindDist(p1, p2 Point) float64 {
	return math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2)
}

// BruteForce finds two points in a given set of 2d points using brute fore approach
// Returns a point and distance between those two points
func BruteForce(pts []Point) ([]Point, float64) {

	dist := math.Inf(1)
	p := make([]Point, 2)

	for i := 0; i < len(pts)-1; i++ {
		for j := i + 1; j < len(pts); j++ {
			d := FindDist(pts[i], pts[j])
			if d < dist {
				dist = d
				p[0] = pts[i]
				p[1] = pts[j]
			}
		}
	}
	return p, dist
}

// ClosestInStrip finds points in a set of points whose distance is less than specified distance
// Returns []Point and distance between them
func ClosestInStrip(pts []Point, dist float64) ([]Point, float64) {
	p := make([]Point, 2)
	d := dist

	for i := 0; i < len(pts); i++ {
		for j := i + 1; j < len(pts) && (pts[j].y-pts[i].y) < d; j++ {
			newDist := FindDist(pts[i], pts[j])
			if newDist < d {
				dist = d
				p[0] = pts[i]
				p[1] = pts[j]
			}
		}
	}

	return p, dist
}

// Ans to sample.txt points
// (8, 1) (10, 7)
// 6.3245
//
// Some other points
//
// 2 3
// 12 30
// 40 50
// 5 1
// 13 31
// 3 4
// -- Ans --
// (2, 3) (3, 4)
// 1.41421
