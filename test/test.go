package main

import "fmt"
import "github.com/dhconnelly/rtreego"

type Thing struct {
	where *rtreego.Rect
	name  string
}

func (t *Thing) Bounds() *rtreego.Rect {
	return t.where
}

func main() {
	fmt.Println("start")
	rt := rtreego.NewTree(2, 3, 3)

	p1 := rtreego.Point{0.4, 0.5}
	//p2 := rtreego.Point{6.2, -3.4}

	r1, _ := rtreego.NewRect(p1, []float64{1, 2})
	//r2, _ := rtreego.NewRect(p2, []float64{1.7, 2.7})

	rt.Insert(&Thing{r1, "foo"})

	//size := rt.Size() // returns 2
	//fmt.Println(size)
}
