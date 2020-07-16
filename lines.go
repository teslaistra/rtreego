package rtreego

import (
	"fmt"
	"reflect"
	"strings"
)

type Line struct {
	name          string
	start, finish Point
	where         *Rect
}

//constructor
func NewLine(p Point, q Point, name string) (r *Line, err error) {
	r = new(Line)
	r.start = p
	r.finish = q
	r.name = name

	r.where, _ = NewRectFromPoints(r.start, r.finish, name)
	return
}

//returns MBR of line. Implementation of interface
func (t *Line) Bounds() *Rect {
	rect := t.where
	return rect
}

//returns type of itself. Implementation of interface
func (t *Line) GetTypeOf() reflect.Type {
	return reflect.TypeOf(t)
}

//returns name of itself. Implementation of interface
func (t *Line) GetNameOf() string {
	return t.name
}

//returns a line from points, stored in rectangle
func (r *Rect) RectToLine(name string) *Line {
	l1, _ := NewLine(r.q, r.p, name)
	return l1
}

//returns length of line
func (l *Line) Length() float64 {
	return GreatCircle(l.start, l.finish)
}

//returns array of points(start and finish point)
func (l *Line) GetPoints() []Point {
	return []Point{l.start, l.finish}
}

//returns string presentation of line
func (r *Line) String() string {
	s := make([]string, 4)
	s[0] = "Start:"
	s[1] = fmt.Sprintf("[%.2f, %.2f]", r.start[0], r.start[1])
	s[2] = fmt.Sprint("End:")
	s[3] = fmt.Sprintf("[%.2f, %.2f]", r.finish[0], r.finish[1])

	return strings.Join(s, " ")
}

//check if lines are equal
func (r *Line) Equal(other *Line) bool {
	for i, e := range r.start {
		if e != other.start[i] {
			return false
		}
	}
	for i, e := range r.finish {
		if e != other.finish[i] {
			return false
		}
	}
	return true
}
