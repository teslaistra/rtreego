package rtreego

import (
	"fmt"
	"reflect"
)

// Point represents a point in n-dimensional Euclidean space.
type Point []float64

//returns MBR of rectangle. Implementation of interface
func (t *Point) Bounds() *Rect {
	rect, _ := NewRectFromPoints(*t, *t, "")
	return rect
}

//returns type of itself. Implementation of interface
func (t *Point) GetTypeOf() reflect.Type {
	return reflect.TypeOf(t)
}

//returns empty string. Implementation of interface
func (t *Point) GetNameOf() string {
	return ""
}

//Returns string representation of rectangle left and right points
func (p *Point) String() string {
	return fmt.Sprintf("%.6f", p)
}

//constructor
func NewPoint(lat, lon float64) (r *Point) {
	r = &Point{lat, lon}
	return r
}
