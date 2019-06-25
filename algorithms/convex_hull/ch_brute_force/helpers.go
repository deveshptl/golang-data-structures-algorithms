package main

import (
	"math"
	"sort"
)

// IsLongestPair checks whether the given first two set of points (p1, p2) have longest distance or not
// The main idea is not to select all given points in question that lies on the same line
func IsLongestPair(p1, p2, p3 Point) bool {
	d12 := math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2)
	d13 := math.Pow(p1.x-p3.x, 2) + math.Pow(p1.y-p3.y, 2)
	d23 := math.Pow(p2.x-p3.x, 2) + math.Pow(p2.y-p3.y, 2)
	if d12 >= d13 && d12 >= d23 {
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

	center.x = center.x / float64(len(pts))
	center.y = center.y / float64(len(pts))

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
