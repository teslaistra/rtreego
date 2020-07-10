package main

import (
	"fmt"
	"reflect"
)
import "github.com/dhconnelly/rtreego"

type Rect struct {
	where *rtreego.Rect
	name  string
}
type Line struct {
	where *rtreego.Line
	name  string
}

var tol = 0.01

type Somewhere struct {
	location rtreego.Point
	name     string
	wormhole chan int
}

func (t *Rect) Bounds() *rtreego.Rect {
	return t.where
}
func (t *Line) Bounds() *rtreego.Rect {

	rect, _ := rtreego.NewRectFromPoints(t.where.PointCoord(0), t.where.PointCoord(1))
	return rect
}

func (t *Rect) GetTypeOf() reflect.Type {
	return reflect.TypeOf(t)
}
func (t *Line) GetTypeOf() reflect.Type {
	return reflect.TypeOf(t)
}

func (t *Rect) GetNameOf() string {
	return t.name
}
func (t *Line) GetNameOf() string {
	return t.name
}

func main() {
	fmt.Println("start")
	rt := rtreego.NewTree(2, 3, 3)

	//p1 := rtreego.Point{0.4, 0.5}
	//p2 := rtreego.Point{6.2, -3.4}
	//r1, _ := rtreego.NewRect(p1, []float64{1, 2})
	//r2, _ := rtreego.NewRect(p2, []float64{1.7, 2.7})
	//rt.Insert(&Rect{r1, "foo"})

	p2 := rtreego.Point{0, 0}
	p3 := rtreego.Point{1, 1}

	l1, _ := rtreego.NewLine(p2, p3)
	//l2, _ := rtreego.NewLine(rtreego.Point{0,2},rtreego.Point{1,2})

	rt.Insert(&Line{l1, "street1"})
	//rt.Insert(&Line{l2, "street2"})

	q := rtreego.Point{0.000000001, 0.000000001}

	results := rt.NearestNeighbor(q)

	//size := rt.Size() // returns 2
	fmt.Println(results.Bounds().Diag())
}
