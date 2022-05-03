// 6kruskal project main.go
package main

import (
	"fmt"
	"math"
	"sort"
)

var (
	n, m int
	v    []point
	e    []road
	s    float64
)

type point struct {
	x, y, depth int
	parent      *point
}

type road struct {
	len  float64
	u, v *point
}

func Union(x, y *point) {
	if x.depth < y.depth {
		x.parent = y
	} else {
		y.parent = x
		if x.depth == y.depth && x != y {
			x.depth++
		}
	}
}

func Find(x *point) *point {
	if x.parent == x {
		return x
	} else {
		y := Find(x.parent)
		x.parent = y
		return y
	}
}

func SpanningTree() {
	j := 0
	for i := 0; i < m && j < n-1; i++ {
		v := Find(e[i].v)
		u := Find(e[i].u)
		if v != u {
			s += e[i].len
			j++
			Union(u, v)
		}
	}
}

func MST_Kruskal() {
	sort.SliceStable(e, func(i, j int) bool {
		return e[i].len < e[j].len
	})
	SpanningTree()
}

func main() {
	var x, y, k int
	fmt.Scan(&n)
	v = make([]point, n)
	m = (n*n - n) / 2
	e = make([]road, m)
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		v[i].x, v[i].y = x, y
		v[i].parent = &(v[i])
		for j := 0; j < i; j++ {
			e[k].v = &(v[i])
			//fmt.Println((*(e[k].v)).x)
			e[k].u = &(v[j])
			e[k].len = math.Sqrt(float64((v[j].x-x)*(v[j].x-x) + (v[j].y-y)*(v[j].y-y)))
			//fmt.Println(e[k].len)
			k++
		}
	}
	MST_Kruskal()
	fmt.Printf("%.2f", s)
}
