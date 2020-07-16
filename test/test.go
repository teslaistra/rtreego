package main

import (
	"fmt"
	"math"
)
import "github.com/dhconnelly/rtreego"

func RoundTwoSigns(x float64) float64 {
	return math.Round(x*100) / 100
}
func TestLines() {
	rt := rtreego.NewTree(2, 3)

	p2 := rtreego.Point{55.766028, 37.636580}
	p3 := rtreego.Point{55.770174, 37.643440}

	l3, _ := rtreego.NewLine(p2, p3, "street1")
	rt.Insert(l3)

	q := rtreego.Point{0.000000001, 0.000000001}

	results := rt.NearestNeighbor(q)
	//size := rt.Size() // returns 2
	fmt.Println(results.(*rtreego.Line))
	fmt.Println(RoundTwoSigns(results.Bounds().Diagonal()))

	fmt.Println(results.Bounds().RectToLine("result line").String())
}

func main() {
	rt := rtreego.NewTree(2, 3)

	p1 := rtreego.NewPoint(13, 13)
	p2 := rtreego.NewPoint(16, 16)
	p3 := rtreego.NewPoint(117, 18)
	p4 := rtreego.NewPoint(12, 12)
	//r2, _ := rtreego.NewRectFromPoints(rtreego.Point{1, 1}, rtreego.Point{2, 5}, "house2")

	r1, _ := rtreego.NewRectFromPoints(rtreego.Point{0, 0}, rtreego.Point{3, 3}, "house1")

	r1.Size()
	l1, _ := rtreego.NewLine(rtreego.Point{0, 0}, rtreego.Point{0, 6}, "street1")
	l2, _ := rtreego.NewLine(rtreego.Point{0, 0}, rtreego.Point{0, 6}, "street1")
	fmt.Println(l1.Equal(l2))
	//f := p1.TestminMaxDist(l3)

	rt.Insert(r1)
	rt.Insert(l1)
	rt.Insert(p1)
	rt.Insert(p2)
	rt.Insert(p3)

	a := rt.NearestNeighbor(*p4)
	fmt.Println(a)

	bb, _ := rtreego.NewRectFromPoints(*p3, *p4, "")

	// Get a slice of the objects in rt that intersect bb:
	results := rt.SearchIntersect(bb)
	fmt.Println(results[0])
}
