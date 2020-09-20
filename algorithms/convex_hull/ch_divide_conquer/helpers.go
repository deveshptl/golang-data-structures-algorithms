package main

import (
	"math"
	"sort"
)

// FindDist finds distance between two points on a plane
func FindDist(p1, p2 Point) float64 {
	return math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2)
}

// FindSide returns an integer specifying whether the point p3 is on left, right or on the line specified by points p1 and p2.
//
// Returns:
//
// -1 if point p3 is on the left side.
//
// 0 if point p3 is on the line.
//
// 1 if point p3 is on the right side.
func FindSide(p1, p2, p3 Point) int {
	area := p1.x*p2.y + p3.x*p1.y + p2.x*p3.y - (p3.x*p2.y + p2.x*p1.y + p1.x*p3.y)
	if area > 0 {
		return 1
	}
	if area < 0 {
		return -1
	}
	return 0
}

// IsInside checks whether a given point p4 is inside a traingle of points p1, p2, p3 or not.
//
// Returns:
//
// true if it is inside the triangle
//
// false if it otherwise
func IsInside(p1, p2, p3, p4 Point) bool {

	var a0, a1, a2, a3 float64

	a0 = (p1.x*(p2.y-p3.y) + p2.x*(p3.y-p1.y) + p3.x*(p1.y-p2.y)) / 2
	a1 = (p1.x*(p2.y-p4.y) + p2.x*(p4.y-p1.y) + p4.x*(p1.y-p2.y)) / 2
	a2 = (p2.x*(p3.y-p4.y) + p3.x*(p4.y-p2.y) + p4.x*(p2.y-p3.y)) / 2
	a3 = (p1.x*(p3.y-p4.y) + p3.x*(p4.y-p1.y) + p4.x*(p1.y-p3.y)) / 2

	// removing -ve sign if any
	if a0 < 0 {
		a0 = -1 * a0
	}
	if a1 < 0 {
		a1 = -1 * a1
	}
	if a2 < 0 {
		a2 = -1 * a2
	}
	if a3 < 0 {
		a3 = -1 * a3
	}

	if math.Round(a0) == math.Round(a1+a2+a3) {
		return true
	}
	return false
}

// FindAngle finds an angle between lines formed by p1p2 and p2p3 represented as ∠p1p2p3
//
// i.e. angle formed at point p2
func FindAngle(p1, p2, p3 Point) float64 {
	deg1 := int(((math.Atan2(p1.x-p2.x, p1.y-p2.y)*180)/math.Pi)+360) % 360
	deg2 := int(((math.Atan2(p3.x-p2.x, p3.y-p2.y)*180)/math.Pi)+360) % 360
	return float64(deg1 - deg2)
}

// SortConvexPoints sorts points of convex polygon in clockwise fashion
func SortConvexPoints(pts []Point) []Point {

	// initializations
	newPts := make([]Point, 0)
	temp := make(map[float64]Point)
	center := Point{x: 0, y: 0}
	p1 := pts[0]

	// find a point within the convex polygon
	for i := 0; i < len(pts); i++ {
		center.x += pts[i].x
		center.y += pts[i].y
	}

	center.x /= float64(len(pts))
	center.y /= float64(len(pts))

	// find angle of each points formed at center point
	for i := 0; i < len(pts); i++ {
		angle := FindAngle(pts[i], center, p1)
		temp[angle] = pts[i]
	}

	// get all keys (angles)
	keys := make([]float64, 0)

	for i := range temp {
		keys = append(keys, i)
	}

	// sort these keys
	sort.Float64s(keys)

	// using map access each points in sorted fashion using keys
	for i := 0; i < len(keys); i++ {
		newPts = append(newPts, temp[keys[i]])
	}

	return newPts
}
