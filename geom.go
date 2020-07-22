// Copyright 2012 Daniel Connelly.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtreego

import (
	"github.com/golang/geo/s2"
	"math"
)

const (
	EARTH_RADIUS = 6371
)

// DimError represents a failure due to mismatched dimensions.
type DimError struct {
	Expected int
	Actual   int
}

type SearchObject struct {
	Object   Spatial
	Distance float64
}

func (err DimError) Error() string {
	return "rtreego: dimension mismatch"
}

// DistError is an improper distance measurement.  It implements the error
// and is generated when a distance-related assertion fails.
type DistError float64

func (err DistError) Error() string {
	return "rtreego: improper distance"
}

// Dist computes the Great Circle distance between two points p and q.
func (p Point) dist(q Point) float64 {
	return GreatCircle(p, q)
}

// minDist computes the square of the distance from a point to a rectangle.
// If the point is contained in the rectangle then the distance is zero.
//
// Implemented per Definition 2 of "Nearest Neighbor Queries" by
// N. Roussopoulos, S. Kelley and F. Vincent, ACM SIGMOD, pages 71-79, 1995.
// http://www.cse.msu.edu/~pramanik/teaching/courses/cse880/14f/lectures/5.multimediaIndexing/KNN-Rousapolis/lec.pdf
// https://www.cs.umd.edu/~nick/papers/nncolor.pdf
func (p Point) MinDist(r *Rect) float64 {
	if len(p) != len(r.p) {
		panic(DimError{len(p), len(r.p)})
	}

	coordinates := make([]float64, 2)

	for i, pi := range p {
		if pi < r.p[i] {
			coordinates[i] = r.p[i]
		} else if pi > r.q[i] {
			coordinates[i] = r.q[i]
		} else {
			coordinates[i] = pi

		}
	}
	return math.Pow(GreatCircle(p, Point{coordinates[0], coordinates[1]}), 2)
}

// minMaxDist computes the minimum of the maximum distances from p to points
// on r.  If r is the bounding box of some geometric objects, then there is
// at least one object contained in r within minMaxDist(p, r) of p.
//
// Implemented per Definition 4 of "Nearest Neighbor Queries" by
// N. Roussopoulos, S. Kelley and F. Vincent, ACM SIGMOD, pages 71-79, 1995.
// http://www.cse.msu.edu/~pramanik/teaching/courses/cse880/14f/lectures/5.multimediaIndexing/KNN-Rousapolis/lec.pdf
// https://www.cs.umd.edu/~nick/papers/nncolor.pdf
func (p Point) minMaxDist(r *Rect) float64 {
	if len(p) != len(r.p) {
		panic(DimError{len(p), len(r.p)})
	}

	RectPoints := []Point{r.p, r.q, {r.p[0], r.q[1]}, {r.q[0], r.p[1]}}
	RectPointsDistances := []float64{
		GreatCircle(p, r.p),                   // to left bottom
		GreatCircle(p, Point{r.p[0], r.q[1]}), // to left top
		GreatCircle(p, r.q),                   // to right top
		GreatCircle(p, Point{r.q[0], r.p[1]}), // to right bottom
		GreatCircle(p, r.p)}                   // to left bottom

	min := math.MaxFloat64
	for k := range RectPoints {
		d := Max(RectPointsDistances[k], RectPointsDistances[k+1])
		if d < min {
			min = d
		}
	}

	return min
}

//Returns distnace in meters
func GreatCircle(from Point, to Point) float64 {
	dLat := (from[0] - to[0]) * (math.Pi / 180.0)
	dLon := (from[1] - to[1]) * (math.Pi / 180.0)

	lat1 := from[0] * (math.Pi / 180.0)
	lat2 := to[0] * (math.Pi / 180.0)

	a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
	a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)

	a := a1 + a2

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EARTH_RADIUS * c * 1000
}

// Size computes the measure of a rectangle (the product of its side lengths).
func (r *Rect) Size() float64 {
	p1 := s2.LatLngFromDegrees(r.p[0], r.p[1])
	p2 := s2.LatLngFromDegrees(r.q[0], r.q[1])
	rect := s2.RectFromLatLng(p1)
	rect = rect.AddPoint(p2)
	return rect.Area()
}

// margin computes the sum of the edge lengths of a rectangle.
func (r *Rect) margin() float64 {
	l1 := GreatCircle(r.p, Point{r.p[0], r.q[1]})
	l2 := GreatCircle(r.q, Point{r.p[0], r.q[1]})

	return (l2 + l1) * 2
}

// containsPoint tests whether p is located inside or on the boundary of r.
func (r *Rect) containsPoint(p Point) bool {
	if len(p) != len(r.p) {
		panic(DimError{len(r.p), len(p)})
	}

	for i, a := range p {
		// p is contained in (or on) r if and only if p <= a <= q for
		// every dimension.
		if a < r.p[i] || a > r.q[i] {
			return false
		}
	}

	return true
}

// containsRect tests whether r2 is is located inside r1.
func (r *Rect) ContainsRect(r2 *Rect) bool {
	if len(r.p) != len(r2.p) {
		panic(DimError{len(r.p), len(r2.p)})
	}

	for i, a1 := range r.p {
		b1, a2, b2 := r.q[i], r2.p[i], r2.q[i]
		// enforced by constructor: a1 <= b1 and a2 <= b2.
		// so containment holds if and only if a1 <= a2 <= b2 <= b1
		// for every dimension.
		if a1 > a2 || b2 > b1 {
			return false
		}
	}

	return true
}

// intersect computes the intersection of two rectangles.  If no intersection
// exists, the intersection is nil.
func Intersect(r1, r2 *Rect) bool {
	dim := len(r1.p)
	if len(r2.p) != dim {
		panic(DimError{dim, len(r2.p)})
	}

	// There are four cases of overlap:
	//
	//     1.  a1------------b1
	//              a2------------b2
	//              p--------q
	//
	//     2.       a1------------b1
	//         a2------------b2
	//              p--------q
	//
	//     3.  a1-----------------b1
	//              a2-------b2
	//              p--------q
	//
	//     4.       a1-------b1
	//         a2-----------------b2
	//              p--------q
	//
	// Thus there are only two cases of non-overlap:
	//
	//     1. a1------b1
	//                    a2------b2
	//
	//     2.             a1------b1
	//        a2------b2
	//
	// Enforced by constructor: a1 <= b1 and a2 <= b2.  So we can just
	// check the endpoints.

	for i := range r1.p {
		a1, b1, a2, b2 := r1.p[i], r1.q[i], r2.p[i], r2.q[i]
		if b2 <= a1 || b1 <= a2 {
			return false
		}
	}
	return true
}

// boundingBox constructs the smallest rectangle containing both r1 and r2.
func boundingBox(r1, r2 *Rect) (bb *Rect) {
	bb = new(Rect)
	dim := len(r1.p)
	bb.p = make([]float64, dim)
	bb.q = make([]float64, dim)
	if len(r2.p) != dim {
		panic(DimError{dim, len(r2.p)})
	}
	for i := 0; i < dim; i++ {
		if r1.p[i] <= r2.p[i] {
			bb.p[i] = r1.p[i]
		} else {
			bb.p[i] = r2.p[i]
		}
		if r1.q[i] <= r2.q[i] {
			bb.q[i] = r2.q[i]
		} else {
			bb.q[i] = r1.q[i]
		}
	}
	return
}

// boundingBoxN constructs the smallest rectangle containing all of r...
func boundingBoxN(rects ...*Rect) (bb *Rect) {
	if len(rects) == 1 {
		bb = rects[0]
		return
	}
	bb = boundingBox(rects[0], rects[1])
	for _, rect := range rects[2:] {
		bb = boundingBox(bb, rect)
	}
	return
}

// Max returns the larger of x or y.
func Max(x, y float64) float64 {
	if x < y {
		return y
	}
	return x
}

//returns shortest distance from point to line using S2
func DistancePointToLine(point Point, r Rect) float64 {
	//A is p of rect, and B is q of rect.
	//C is a point

	//line
	ll1 := s2.LatLngFromDegrees(r.p[0], r.p[1])
	ll2 := s2.LatLngFromDegrees(r.q[0], r.q[1])
	pointA := s2.PointFromLatLng(ll1)
	pointB := s2.PointFromLatLng(ll2)

	ll3 := s2.LatLngFromDegrees(point[0], point[1])
	pointC := s2.PointFromLatLng(ll3)

	EarthRadius := 6370986.884258304
	dist := s2.DistanceFromSegment(pointC, pointA, pointB).Radians() * EarthRadius
	return dist
}

//returns shortest distance between Rectangle and line
func DistanceRectToLine(rect Rect, line Line) float64 {

	var distances []float64

	distances = append(distances, DistancePointToLine(rect.q, *line.Bounds()))
	distances = append(distances, DistancePointToLine(rect.p, *line.Bounds()))
	rectangle_point2 := Point{rect.p[0], rect.q[1]}
	rectangle_point4 := Point{rect.q[0], rect.p[1]}
	distances = append(distances, DistancePointToLine(rectangle_point2, *line.Bounds()))
	distances = append(distances, DistancePointToLine(rectangle_point4, *line.Bounds()))

	side, _ := NewLine(rect.p, rectangle_point2, "")
	side1 := *side.Bounds()
	side, _ = NewLine(rect.p, rectangle_point4, "")
	side2 := *side.Bounds()
	side, _ = NewLine(rect.q, rectangle_point2, "")
	side3 := *side.Bounds()
	side, _ = NewLine(rect.q, rectangle_point4, "")
	side4 := *side.Bounds()

	distances = append(distances, DistancePointToLine(line.start, side1))
	distances = append(distances, DistancePointToLine(line.start, side2))
	distances = append(distances, DistancePointToLine(line.start, side3))
	distances = append(distances, DistancePointToLine(line.start, side4))

	distances = append(distances, DistancePointToLine(line.finish, side1))
	distances = append(distances, DistancePointToLine(line.finish, side2))
	distances = append(distances, DistancePointToLine(line.finish, side3))
	distances = append(distances, DistancePointToLine(line.finish, side4))

	min := distances[0]
	for _, v := range distances {
		if v < min {
			min = v
		}
	}
	return min
}

//returns shortest distance between two lines
func DistanceLineToLine(from Line, to Line) float64 {
	var distances []float64

	distances = append(distances, DistancePointToLine(from.start, *to.Bounds()))
	distances = append(distances, DistancePointToLine(from.finish, *to.Bounds()))

	min := distances[0]
	for _, v := range distances {
		if v < min {
			min = v
		}
	}
	return min
}
