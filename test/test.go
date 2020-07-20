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

	p1 := rtreego.NewPoint(55.753230, 37.622516)

	p2 := rtreego.NewPoint(55.753241, 37.621807)
	p3 := rtreego.NewPoint(55.753097, 37.619399)
	p4 := rtreego.NewPoint(55.753188, 37.618384)
	p5 := rtreego.NewPoint(55.753193, 37.617537)

	p6 := rtreego.NewPoint(55.753199, 37.617327)
	p7 := rtreego.NewPoint(55.753272, 37.616517)
	p8 := rtreego.NewPoint(55.753306, 37.616096)
	p9 := rtreego.NewPoint(55.753336, 37.615610)
	p10 := rtreego.NewPoint(55.753361, 37.614410)
	p11 := rtreego.NewPoint(55.753330, 37.613546)
	p12 := rtreego.NewPoint(55.753318, 37.613222)
	p13 := rtreego.NewPoint(55.753294, 37.613028)

	//r2, _ := rtreego.NewRectFromPoints(rtreego.Point{0,0}, rtreego.Point{1,0}, "house2")

	//r1, _ := rtreego.NewRectFromPoints(rtreego.Point{55.749748, 37.623296}, rtreego.Point{55.749449, 37.623317}, "house1")

	//r1.Size()
	l1, _ := rtreego.NewLine(rtreego.Point{55.747051, 37.608887}, rtreego.Point{55.745990, 37.609685}, "street1")
	//	l2, _ := rtreego.NewLine(rtreego.Point{0, 0}, rtreego.Point{0, 6}, "street1")
	//f := p1.TestminMaxDist(l3)

	//rt.Insert(r2)
	//rt.Insert(l1)
	rt.Insert(p1)
	rt.Insert(p2)
	rt.Insert(p3)
	rt.Insert(p4)
	rt.Insert(p5)
	rt.Insert(p6)
	rt.Insert(p7)
	rt.Insert(p8)
	rt.Insert(p9)
	rt.Insert(p10)
	rt.Insert(p11)
	rt.Insert(p12)
	rt.Insert(p13)

	// Get a slice of the objects in rt that intersect bb:
	fmt.Println(rt.NnInRadiusPoint(11, 1000, *p1))
	fmt.Println("kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk")
	fmt.Println(rt.NnInRadiusLine(11, 1000, *l1))
}
