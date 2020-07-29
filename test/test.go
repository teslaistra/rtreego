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
	fmt.Println(rt.NnInRadiusLine(100, 120, *arbat, ""))
	fmt.Println("____________sort________________")
	//	rt.Delete(result[2])
	fmt.Println(rt.NnInRadiusLine(100, 120, *arbat, "sort"))

}
func stressPoints(rt rtreego.Rtree) {

	num := 10000000
	fmt.Println("Загружаю", num, "точек")
	start := time.Now()

	for i := 0; i < num; i++ {
		lat := randFloat(55.765654, 55.755966)
		lon := randFloat(37.643501, 37.583127)
		rt.Insert(rtreego.NewPoint(lat, lon))
	}
	elapsed := time.Since(start)
	fmt.Println("Inserting took", elapsed)
	maxDist := 60000

	for i := 0; i <= maxDist; i += 5000 {
		start := time.Now()
		lat := randFloat(55.765654, 55.755966)
		lon := randFloat(37.643501, 37.583127)
		fmt.Println(rt.NnInRadiusPoint(10, float64(i), *rtreego.NewPoint(lat, lon), "sort"))
		elapsed := time.Since(start)
		fmt.Println("Searching in radius", i, "took via Sort", elapsed)

		start = time.Now()
		fmt.Println(rt.NnInRadiusPoint(10, float64(i), *rtreego.NewPoint(lat, lon), ""))
		fmt.Println("Searching in radius", i, "took via Quicksort", time.Since(start))
		fmt.Println("At point", lat, lon)
		fmt.Println("_________________________________________________________")

	}
}

func stressLines(rt rtreego.Rtree) {

	num := 100000
	fmt.Println("Inserting", num, "lines")
	start := time.Now()

	for i := 0; i < num; i++ {
		latA := randFloat(55.765654, 55.755966)
		lonA := randFloat(37.643501, 37.583127)
		latB := randFloat(55.765654, 55.755966)
		lonB := randFloat(37.643501, 37.583127)
		line, _ := rtreego.NewLine(rtreego.Point{latA, lonA}, rtreego.Point{latB, lonB}, "")
		rt.Insert(line)
	}
	elapsed := time.Since(start)
	fmt.Println("Inserting took", elapsed)
	maxDist := 60000

	for i := 0; i <= maxDist; i += 5000 {
		start := time.Now()
		latA := randFloat(55.765654, 55.755966)
		lonA := randFloat(37.643501, 37.583127)
		latB := randFloat(55.765654, 55.755966)
		lonB := randFloat(37.643501, 37.583127)
		line, _ := rtreego.NewLine(rtreego.Point{latA, lonA}, rtreego.Point{latB, lonB}, "")
		fmt.Println(rt.NnInRadiusLine(10, float64(i), *line, "sort"))
		elapsed := time.Since(start)
		fmt.Println("Searching in radius", i, "took via Sort", elapsed)

		start = time.Now()
		fmt.Println(rt.NnInRadiusLine(10, float64(i), *line, ""))
		fmt.Println("Searching in radius", i, "took via Quicksort", time.Since(start))
		fmt.Println("At line", latA, lonA, "  --  ", latB, lonB)
		fmt.Println("_________________________________________________________")

	}
}
func main() {

	//rt := rtreego.NewTree(10, 20)

	//p1 := rtreego.NewPoint(55.752612, 37.581785)
	//p2 := rtreego.NewPoint(55.752575, 37.580905)

	//rt.Insert(p1)
	//rt.Insert(p2)

	//fmt.Println(rt.NnInRadiusPoint(1, 2000, *rtreego.NewPoint(55.752588, 37.578526), "sort"))
	//test1(*rt)
	//stress(*rt)
	//stressLines(*rt)
	from, _ := rtreego.NewLine(rtreego.Point{55.687898, 37.525656}, rtreego.Point{55.685431, 37.520978}, "")
	to, _ := rtreego.NewLine(rtreego.Point{55.688026, 37.520583}, rtreego.Point{55.684284, 37.526417}, "")

	fmt.Println(rtreego.DistanceLineToLine(*from, *to))

}

func randFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano() * rand.Int63())

	res := min + rand.Float64()*(max-min)
	return res
}
