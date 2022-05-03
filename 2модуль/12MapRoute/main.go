package main

import (
	"fmt"
)

var (
	n int
	a [][1501]Point
	q queue
)

type Point struct {
	dist, index int
	l, i, j     int
}

type queue struct {
	heap       [](*Point)
	cap, count int
}

func Insert(ptr *Point) {
	i := q.count
	q.count++
	q.heap[i] = ptr
	for i > 0 && (q.heap[(i-1)/2].dist > q.heap[i].dist) {
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
		if l < n && q.heap[i].dist > q.heap[l].dist {
			i = l
		}
		if r < n && q.heap[i].dist > q.heap[r].dist {
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
	ptr.dist = k
	for i > 0 && q.heap[(i-1)/2].dist > k {
		t := q.heap[(i-1)/2]
		q.heap[(i-1)/2] = q.heap[i]
		q.heap[i] = t
		q.heap[i].index = i
		i = (i - 1) / 2
	}
	ptr.index = i
}

func Relax(u, v *Point) bool {
	changed := u.dist+v.l < v.dist
	if changed {
		v.dist = u.dist + v.l
	}
	return changed
}

func Dijkstra(s *Point) {
	q.heap = make([]*Point, n*n)
	q.cap = n * n
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			v := &a[i][j]
			if v != s {
				v.dist = 100000
			}
			Insert(v)
		}
	}
	for q.count != 0 {
		v := ExtractMin()
		v.index = -1
		i := v.i
		j := v.j
		if i > 0 {
			u := &a[i-1][j]
			if u.index != -1 && Relax(v, u) {
				DecreaseKey(u, u.dist)
			}
		}
		if i < n-1 {
			u := &a[i+1][j]
			if u.index != -1 && Relax(v, u) {
				DecreaseKey(u, u.dist)
			}
		}
		if j > 0 {
			u := &a[i][j-1]
			if u.index != -1 && Relax(v, u) {
				DecreaseKey(u, u.dist)
			}
		}
		if j < n-1 {
			u := &a[i][j+1]
			if u.index != -1 && Relax(v, u) {
				DecreaseKey(u, u.dist)
			}
		}
	}
}

func main() {
	fmt.Scan(&n)
	var a0 [1501][1501]Point
	a = a0[:n][:n]
	l := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&l)
			a[i][j].l = l
			a[i][j].i = i
			a[i][j].j = j
		}
	}
	a[0][0].dist = a[0][0].l
	Dijkstra(&a[0][0])
	fmt.Println(a[n-1][n-1].dist)
}
