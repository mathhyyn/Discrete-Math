package main

import (
	"fmt"
)

var g []Point
var N, M, K, A int
var f, error, b bool

type Elem struct {
	next *Elem
	v    int
}

type Point struct {
	next *Elem
	used bool
	dist []int
}

func component(x *Elem, k int) {
	k++
	for x != nil {
		if x.v != A && g[A].dist[x.v] == 0 || g[A].dist[x.v] > k || g[x.v].used == f {
			g[A].dist[x.v] = k
			g[x.v].used = !f
			component(g[x.v].next, k)
		}
		x = x.next
	}
}

func main() {
	var (
		d0         [1000001]Point
		a0         [2000001]Elem
		r          [100001]int
		x, y, i, j int
	)
	fmt.Scan(&N)
	fmt.Scan(&M)
	g = d0[:N]
	a := a0[:2*M]
	j = 0
	for i = 0; i < M; i++ {
		fmt.Scan(&x, &y)
		if g[x].next == nil {
			a[j].v = y
			g[x].next = &a[j]
		} else {
			a[j].v = y
			t := g[x].next
			g[x].next = &a[j]
			a[j].next = t
		}
		j++
		if g[y].next == nil {
			a[j].v = x
			g[y].next = &a[j]
		} else {
			a[j].v = x
			t := g[y].next
			g[y].next = &a[j]
			a[j].next = t
		}
		j++
	}
	fmt.Scan(&K)
	roots := r[:K]
	for i = 0; i < K; i++ {
		fmt.Scan(&x)
		roots[i] = x
		var d0 [1000001]int
		d := d0[:N]
		g[x].dist = d
		A = x
		//fmt.Println(f, g[x].used)
		if f != g[x].used {
			fmt.Print("-")
			error = true
			break
		}
		g[x].used = !f
		component(g[x].next, 0)
		f = !f
	}

	if !error {
		for i = 0; i < N; i++ {
			x = g[roots[0]].dist[i]
			for j = 1; j < K; j++ {
				y = g[roots[j]].dist[i]
				//fmt.Println(g[roots[j]].dist[i], g[roots[j-1]].dist[i])
				if x == 0 || x != y {
					error = true
					break
				}
				x = y
			}
			if !error {
				fmt.Print(i, " ")
				b = true
			}
			error = false
		}
		if !b {
			fmt.Print("-")
		}

	}
}
