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

//конструктор
func NewLine(p Point, q Point, name string) (r *Line, err error) {
	r = new(Line)
	r.start = p
	r.finish = q
	r.name = name

	r.where, _ = NewRectFromPoints(r.start, r.finish, name)
	return
}

func (t *Line) Bounds() *Rect {
	rect := t.where
	return rect
}

func (t *Line) GetTypeOf() reflect.Type {
	return reflect.TypeOf(t)
}

func (t *Line) GetNameOf() string {
	return t.name
}

func (r *Rect) RectToLine(name string) *Line {
	l1, _ := NewLine(r.q, r.p, name)
	return l1
}

func (l *Line) Lenght() float64 {
	return GreatCircle(l.start, l.finish)
}

func (l *Line) GetPoints() []Point {
	return []Point{l.start, l.finish}
}

func (r *Line) String() string {
	s := make([]string, 4)
	s[0] = "Start:"
	s[1] = fmt.Sprintf("[%.2f, %.2f]", r.start[0], r.start[1])
	s[2] = fmt.Sprint("End:")
	s[3] = fmt.Sprintf("[%.2f, %.2f]", r.finish[0], r.finish[1])

	return strings.Join(s, " ")
}

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
