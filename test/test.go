package main

import (
	"fmt"
	"math"
)
import "github.com/dhconnelly/rtreego"

func RoundTwoSigns(x float64) float64 {
	return math.Round(x*100) / 100
}

func main() {
	rt := rtreego.NewTree(2, 3)

	arbat, _ := rtreego.NewLine(rtreego.Point{55.752575, 37.575047}, rtreego.Point{55.752624, 37.582622}, "arbat")
	p1 := rtreego.NewPoint(55.752612, 37.581785)
	p2 := rtreego.NewPoint(55.752575, 37.580905)
	out := rtreego.NewPoint(55.753665, 37.580948)

	rt.Insert(p1)
	rt.Insert(p2)
	rt.Insert(out)

	fmt.Println(rt.NnInRadiusLine(100, 120, *arbat))
}
