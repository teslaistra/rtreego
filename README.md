rtreego
=======

A fork of library for efficiently storing and querying geospatial data
in the Go programming language.

About
-----

The R-tree is a popular data structure for efficiently storing and
querying spatial objects; one common use is implementing geospatial
indexes in database management systems. Both bounding-box queries
and k-nearest-neighbor queries are supported.

R-trees are balanced, so maximum tree height is guaranteed to be
logarithmic in the number of entries; however, good worst-case
performance is not guaranteed.  Instead, a number of rebalancing
heuristics are applied that perform well in practice.  For more
details please refer to the references.

This implementation is done for working with geospatial data.
All distances are calculated with Great Circle formula.

Getting Started
---------------

Get the source code from [GitHub](https://github.com/teslaistra/rtreego) or,
with Go 1 installed, run `https://github.com/teslaistra/rtreego`.

Make sure you `import github.com/teslaistra/rtreego` in your Go source files.

Documentation
-------------

### Storing, updating, and deleting objects

To create a new tree specify the minimum
and maximum branching factor:

    rt := rtreego.NewTree(25, 50)

You can also bulk-load the tree when creating it by passing the objects as
a parameter.

    rt := rtreego.NewTree(25, 50, objects...)

Any type that implements the `Spatial` interface can be stored in the tree:

    type Spatial interface {
      Bounds() *Rect
      GetTypeOf() reflect.Type
      GetNameOf() string
    }

`Rect`s are data structures for representing spatial objects, while `Point`s
represent spatial locations.  Creating `Point`s is easy--they're just slices
of `float64`s:

    p1 := rtreego.Point{0.4, 0.5}
    p2 := rtreego.Point{6.2, -3.4}
or
    p3 := rtreego.NewPoint(3, 6)

To create a `Rect`, specify left bootom point, and right top point, give it a name:

    r1, _ := rtreego.NewRectFromPoints(p1, rtreego.Point{1, 2}, "rectangle 1")
    r2, _ := rtreego.NewRectFromPoints(p2, p1, "rect2")

To demonstrate creating your own shape, let's create and store some test data. You can create every type you want, but it should be possible to create a Minimal Bounding Rectangle(MBR) around it. 

    type Thing struct {
      where *Rect
      name string
    }

    func (t *Thing) Bounds() *Rect {
      return t.where
    }
    
    func (t *Thing) GetTypeOf() reflect.Type {
	    return reflect.TypeOf(t)
    }
    
    func (t *Thing) GetNameOf() string {
	    return t.name
    }
    
    rt.Insert(&Thing{r1, "foo"})
    rt.Insert(&Thing{r2, "bar"})

    size := rt.Size() // returns 2

We can insert and delete objects from the tree in any order.

    rt.Delete(thing2)
    // do some stuff...
    rt.Insert(anotherThing)

<b>Comparator functionality is implemented by original lib, and it needs to be tested</b><br>

Note that ```Delete``` function does the equality comparison by comparing the
memory addresses of the objects. If you do not have a pointer to the original
object anymore, you can define a custom comparator.

    type Comparator func(obj1, obj2 Spatial) (equal bool)

You can use a custom comparator with ```DeleteWithComparator``` function.

    cmp := func(obj1, obj2 Spatial) bool {
      sp1 := obj1.(*IDRect)
      sp2 := obj2.(*IDRect)

      return sp1.ID == sp2.ID
    }

    rt.DeleteWithComparator(obj, cmp)


If you want to update the location of an object, you must delete it, update it,
and re-insert.  Just modifying the object so that the `*Rect` returned by
`Location()` changes, without deleting and re-inserting the object, will
corrupt the tree.

### Queries

Bounding-box and k-nearest-neighbors near point, k-nearest-neighbors in radius near line and point queries are supported.

Bounding-box queries require a search `*Rect`. It returns all objects that
touch the search rectangle.

    bb, _ := rtreego.NewRect(rtreego.Point{1.7, -3.4}, []float64{3.2, 1.9})

    // Get a slice of the objects in rt that intersect bb:
    results := rt.SearchIntersect(bb)

### Filters
<b>Filter functionality is implemented by original lib, and it needs to be tested</b><br>

You can filter out values during searches by implementing Filter functions.

    type Filter func(results []Spatial, object Spatial) (refuse, abort bool)

A filter for limiting results by result count is included in the package for
backwards compatibility.

    // maximum of three results will be returned
    tree.SearchIntersect(bb, LimitFilter(3))

Nearest-neighbor queries find the objects in a tree closest to a specified
query point.

    q := rtreego.Point{6.5, -2.47}
    k := 5

    // Get a slice of the k objects in rt closest to q:
    results = rt.NearestNeighbors(k, q)

### More information

See [GoDoc](http://godoc.org/github.com/dhconnelly/rtreego) for full API
documentation.

References
----------

- A. Guttman.  R-trees: A Dynamic Index Structure for Spatial Searching.
  Proceedings of ACM SIGMOD, pages 47-57, 1984.
  http://www.cs.jhu.edu/~misha/ReadingSeminar/Papers/Guttman84.pdf

- N. Beckmann, H .P. Kriegel, R. Schneider and B. Seeger.  The R*-tree: An
  Efficient and Robust Access Method for Points and Rectangles.  Proceedings
  of ACM SIGMOD, pages 323-331, May 1990.
  http://infolab.usc.edu/csci587/Fall2011/papers/p322-beckmann.pdf

- N. Roussopoulos, S. Kelley and F. Vincent.  Nearest Neighbor Queries.  ACM
  SIGMOD, pages 71-79, 1995.
  http://www.postgis.org/support/nearestneighbor.pdf
  
- Nick Roussopoulos, Stephen Kelley, Frederic Vincent, University of Maryland, May 1995, Nearest Neighbor Queries 
  https://www.cs.umd.edu/~nick/papers/nncolor.pdf
  
- Advanced Database Systems, KNN- Search, Roussopoulos Paper, S. Pramanik
  http://www.cse.msu.edu/~pramanik/teaching/courses/cse880/14f/lectures/5.multimediaIndexing/KNN-Rousapolis/lec.pdf


Author
------
Written by [Daniel Connelly](http://dhconnelly.com) (<dhconnelly@gmail.com>).
Edited by [Daniil_Yefimov](https://github.com/teslaistra) (danyefimoff@gmail.com)

License
-------

rtreego is released under a BSD-style license, described in the `LICENSE`
file.
