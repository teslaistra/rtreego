package main

import (
	"fmt"
	"math"
)
import "github.com/dhconnelly/rtreego"

type Rect struct {
	where *rtreego.Rect
	name  string
}

var tol = 0.01

type Somewhere struct {
	location rtreego.Point
	name     string
	wormhole chan int
}

func RoundTwoSigns(x float64) float64 {
	return math.Round(x*100) / 100
}
func TestLines() {
	rt := rtreego.NewTree(2, 3, 3)

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
	rt := rtreego.NewTree(2, 3, 3)

	//p1 := rtreego.Point{3, 1}
	r1, _ := rtreego.NewRectFromPoints(rtreego.Point{0, 0}, rtreego.Point{0, 6}, "house1")
	l1, _ := rtreego.NewLine(rtreego.Point{0, 0}, rtreego.Point{0, 6}, "street1")

	//f := p1.TestminMaxDist(l3)

	rt.Insert(r1)
	rt.Insert(l1)

	//fmt.Println(f)

}
