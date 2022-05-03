// 6kruskal project main.go
package main

import (
	"fmt"
)

var (
	n, m, s int
	g       []Point
	q       queue
	v       *Point
)

type queue struct {
	heap       [](*Point)
	cap, count int
}

type Elem struct {
	next   *Elem
	v, len int
}

type Point struct {
	next       *Elem
	index, key int
	value      *Point
}

func MST_Prim() {
	for {
		v.index = -2
		x := v.next
		for x != nil {
			u := &g[x.v]
			if u.index == -1 {
				u.key = x.len
				u.value = v
				Insert(u)
			} else {
				if u.index != -2 && u.key > x.len {
					u.value = v
					DecreaseKey(u, x.len)
				}
			}
			x = x.next
		}
		if q.count == 0 {
			break
		}
		v = ExtractMin()
		s += v.key

	}
}

func Insert(ptr *Point) {
	i := q.count
	q.count++
	q.heap[i] = ptr
	for i > 0 && (q.heap[(i-1)/2].key > q.heap[i].key) {
		t := q.heap[(i-1)/2]
		q.heap[(i-1)/2] = q.heap[i]
		q.heap[i] = t
		q.heap[i].index = i
		i = (i - 1) / 2
	}
	q.heap[i].index = i
}

func ExtractMin() *Point {
	ptr := q.heap[0]
	q.count--
	if q.count > 0 {
		q.heap[0] = q.heap[q.count]
		q.heap[0].index = 0
		Heapify(q.count)
	}
	return ptr
}

func Heapify(n int) {
	i := 0
	for {
		l := 2*i + 1
		r := l + 1
		j := i
		if l < n && q.heap[i].key > q.heap[l].key {
			i = l
		}
		if r < n && q.heap[i].key > q.heap[r].key {
			i = r
		}
		if i == j {
			break
		}
		t := q.heap[i]
		q.heap[i] = q.heap[j]
		q.heap[j] = t
		q.heap[i].index = i
		q.heap[j].index = j
	}
}

func DecreaseKey(ptr *Point, k int) {
	i := ptr.index
	ptr.key = k
	for i > 0 && q.heap[(i-1)/2].key > k {
		t := q.heap[(i-1)/2]
		q.heap[(i-1)/2] = q.heap[i]
		q.heap[i] = t
		q.heap[i].index = i
		i = (i - 1) / 2
	}
	ptr.index = i
}

func main() {
	var (
		x, y, len int
		g0        [1000001]Point
		a0        [2000001]Elem
		h         [1000001]*Point
	)
	fmt.Scan(&n)
	fmt.Scan(&m)
	q.heap = h[:n]
	q.cap = n
	g = g0[:n]
	a := a0[:2*m]
	j := 0
	for i := 0; i < m; i++ {
		fmt.Scan(&x, &y, &len)
		if g[x].next == nil {
			a[j].v = y
			g[x].next = &a[j]
			g[x].index = -1
		} else {
			a[j].v = y
			t := g[x].next
			g[x].next = &a[j]
			a[j].next = t
		}
		a[j].len = len
		j++
		if g[y].next == nil {
			a[j].v = x
			g[y].next = &a[j]
			g[y].index = -1
		} else {
			a[j].v = x
			t := g[y].next
			g[y].next = &a[j]
			a[j].next = t
		}
		a[j].len = len
		j++
	}
	v = &g[0]
	MST_Prim()
	fmt.Println(s)
}
