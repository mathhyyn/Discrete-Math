package main

import (
	"fmt"
)

var (
	d            []Point
	n, m, bridge int
	component    int
	q            Queue
)

type Queue struct {
	data                   []int
	cap, count, head, tail int
}

type Elem struct {
	next *Elem
	v    int
}

type Point struct {
	next         *Elem
	comp, parent int
	used         bool
	mark         byte
}

func Enqueue(x int) {
	q.data[q.tail] = x
	q.tail++
	if q.tail == q.cap {
		q.tail = 0
	}
	q.count++
}

func Dequeue() int {
	x := q.data[q.head]
	q.head++
	if q.head == q.cap {
		q.head = 0
	}
	q.count--
	return x
}

func VisitVertex1(v int) {
	d[v].mark = 1
	Enqueue(v)
	b := d[v].next
	for b != nil {
		u := b.v
		if d[u].mark == 0 {
			d[u].parent = v
			VisitVertex1(u)
		}
		b = b.next
	}
	d[v].mark = 2
}

func VisitVertex2(v int) {
	d[v].comp = component
	b := d[v].next
	for b != nil {
		u := b.v
		if d[u].comp < component && d[u].comp != -1 {
			bridge++
		}
		if d[u].comp == -1 && d[u].parent != v {
			VisitVertex2(u)
		}
		b = b.next
	}
}

func main() {
	var (
		d0         [1000001]Point
		a0         [2000001]Elem
		q0         [1000001]int
		x, y, i, j int
	)

	fmt.Scan(&n)
	fmt.Scan(&m)
	q.data = q0[:n]
	q.cap = n
	d = d0[:n]
	a := a0[:2*m]
	j = 0
	for i = 0; i < m; i++ {
		fmt.Scan(&x, &y)
		if d[x].next == nil {
			a[j].v = y
			d[x].next = &a[j]
		} else {
			a[j].v = y
			t := d[x].next
			d[x].next = &a[j]
			a[j].next = t
		}
		j++
		if d[y].next == nil {
			a[j].v = x
			d[y].next = &a[j]
		} else {
			a[j].v = x
			t := d[y].next
			d[y].next = &a[j]
			a[j].next = t
		}
		j++
	}

	for i = 0; i < n; i++ {
		if d[i].mark == 0 {
			VisitVertex1(i)
		}
	}

	for i = 0; i < n; i++ {
		d[i].comp = -1
	}
	for q.count != 0 {
		v := Dequeue()
		if d[v].comp == -1 {
			VisitVertex2(v)
			component++
		}
	}

	fmt.Println(bridge)
}
