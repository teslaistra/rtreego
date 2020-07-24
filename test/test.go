package main

import (
	"fmt"
	"math"
	"time"
)
import "math/rand"
import "github.com/dhconnelly/rtreego"

func RoundTwoSigns(x float64) float64 {
	return math.Round(x*100) / 100
}

func test1(rt rtreego.Rtree) {
	arbat, _ := rtreego.NewLine(rtreego.Point{55.752575, 37.575047}, rtreego.Point{55.752624, 37.582622}, "arbat")
	fmt.Println(arbat.GetNameOf(), "len:", arbat.Length())
	p1 := rtreego.NewPoint(55.752612, 37.581785)
	p2 := rtreego.NewPoint(55.752575, 37.580905)
	out := rtreego.NewPoint(55.753665, 37.580948)

	rt.Insert(p1)
	rt.Insert(p2)
	rt.Insert(out)
	result := rt.NnInRadiusLine(100, 120, *arbat)
	fmt.Println(rt.NnInRadiusLine(100, 120, *arbat))

	rt.Delete(result[2])
	fmt.Println(rt.NnInRadiusLine(100, 120, *arbat))

}

func main() {
	rt := rtreego.NewTree(10, 20)
	num := 10000000
	fmt.Println("0Вставляю", num, "точек")
	start := time.Now()

	for i := 0; i < num; i++ {
		lat := randFloat(55.765654, 55.755966)
		lon := randFloat(37.643501, 37.583127)
		//fmt.Println(lat, lon)
		rt.Insert(rtreego.NewPoint(lat, lon))
	}
	elapsed := time.Since(start)
	fmt.Println("Inserting took", elapsed)
	maxDist := 60000

	for i := 0; i <= maxDist; i += 5000 {
		start := time.Now()
		lat := randFloat(55.765654, 55.755966)
		lon := randFloat(37.643501, 37.583127)
		rt.NnInRadiusPoint(10, float64(i), *rtreego.NewPoint(lat, lon))
		elapsed := time.Since(start)
		fmt.Println("Searchig in radius", i, "took", elapsed)

	}
}

func randFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())

	res := min + rand.Float64()*(max-min)
	return res
}
