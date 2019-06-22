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
	fmt.Println("\n-- Closest Pair Problem using Brute Force Approach --")
	p, dist := ClosestPair(arr)
	fmt.Println("\nClosest pair of points are:", p[0], p[1], "whose distance is:", math.Sqrt(dist))
}

// ClosestPair finds two points in a given set of 2d points using BruteForce
// Returns a point and distance between those two points
func ClosestPair(pts []Point) ([]Point, float64) {

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

// FindDist finds distance between two points
func FindDist(p1, p2 Point) float64 {
	return math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2)
}
