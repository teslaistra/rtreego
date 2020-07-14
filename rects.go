package rtreego

import (
	"fmt"
	"reflect"
	"strings"
)

// Rect represents a subset of n-dimensional Euclidean space of the form
// [a1, b1] x [a2, b2] x ... x [an, bn], where ai < bi for all 1 <= i <= n.
type Rect struct {
	p, q Point // Enforced by NewRect: p[i] <= q[i] for all i.
	name string
}

//returns MBR of rectangle. Implementation of interface
func (t *Rect) Bounds() *Rect {
	//rect, _ := NewRectFromPoints(t.p, t.q)
	return t
}

//returns type of itself. Implementation of interface
func (t *Rect) GetTypeOf() reflect.Type {
	return reflect.TypeOf(t)
}

//returns name of itself. Implementation of interface
func (t *Rect) GetNameOf() string {
	return t.name
}

// NewRectFromPoints constructs and returns a pointer to a Rect given a corner points.
func NewRectFromPoints(minPoint, maxPoint Point, name string) (r *Rect, err error) {
	if len(minPoint) != len(maxPoint) {
		err = &DimError{len(minPoint), len(maxPoint)}
		return
	}
	//checking that  min and max points is swapping
	for i, p := range minPoint {
		if minPoint[i] > maxPoint[i] {
			minPoint[i] = maxPoint[i]
			maxPoint[i] = p
		}
	}

	r = &Rect{p: minPoint, q: maxPoint, name: name}

	return
}

//Diagonal returns rectangle diagonal
func (r *Rect) Diagonal() float64 {
	return GreatCircle(r.p, r.q)
}

// Equal returns true if the two rectangles are equal
//TODO ПРОТЕСТИРОВАТЬ
func (r *Rect) Equal(other *Rect) bool {
	for i, e := range r.p {
		if e != other.p[i] {
			return false
		}
	}
	for i, e := range r.q {
		if e != other.q[i] {
			return false
		}
	}
	return true
}

//Returns string representation of rectangle left and right points
func (r *Rect) String() string {
	s := make([]string, 4)
	s[0] = "Left:"
	s[1] = fmt.Sprintf("[%.2f, %.2f]", r.p[0], r.p[1])
	s[2] = fmt.Sprint("Right:")
	s[3] = fmt.Sprintf("[%.2f, %.2f]", r.q[0], r.q[1])

	return strings.Join(s, " ")
}

// NewRect constructs and returns a pointer to a Rect given a corner point and
// the lengths of each dimension.  The point p should be the most-negative point
// on the rectangle (in every dimension) and every length should be positive.
func NewRect(p Point, lengths []float64) (r *Rect, err error) {
	r = new(Rect)
	r.p = p
	if len(p) != len(lengths) {
		err = &DimError{len(p), len(lengths)}
		return
	}
	r.q = make([]float64, len(p))
	for i := range p {
		if lengths[i] <= 0 {
			err = DistError(lengths[i])
			return
		}
		r.q[i] = p[i] + lengths[i]
	}
	return
}
