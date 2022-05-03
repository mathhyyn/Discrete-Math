package main

import (
	"fmt"
)

var d []Point
var n, m int

type Elem struct {
	next *Elem
	v    int
}

type Point struct {
	next *Elem
	k, n int
	used bool
}

func component(x *Elem) (int, int) {
	n, l := 1, 0
	for x != nil {
		if !d[x.v].used {
			d[x.v].used = true
			l += d[x.v].k
			n0, l0 := component(d[x.v].next)
			n += n0
			l += l0
		}
		x = x.next
	}
	return n, l
}

func flag(x *Elem) {
	for x != nil {
		if d[x.v].used {
			d[x.v].used = false
			flag(d[x.v].next)
		}
		x = x.next
	}
}

func main() {
	var (
		d0                 [1000001]Point
		a0                 [2000001]Elem
		x, y, i, j, max, c int
	)
	fmt.Scan(&n)
	fmt.Scan(&m)
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
		d[x].k++
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
		if !d[i].used {
			d[i].used = true
			m, l := component(d[i].next)
			d[i].k += l
			//fmt.Println(i, d[i].k)
			if m > max || (m == max && d[i].k > d[c].k) {
				max = m
				c = i
			}
		}
	}

	//fmt.Println(max, c, d[c].k, d[48].k)
	d[c].used = false
	flag(d[c].next)

	fmt.Println("graph {")
	for i = 0; i < n; i++ {
		fmt.Print(i)
		if !d[i].used {
			fmt.Println(" [color = red]")
			b := d[i].next
			for b != nil && !d[i].used {
				if i <= b.v {
					fmt.Println(i, " -- ", b.v, " [color = red]")
					if i == b.v {
						b = b.next
					}
				}
				b = b.next
			}
		} else {
			fmt.Println()
			b := d[i].next
			for b != nil {
				if i <= b.v {
					fmt.Println(i, " -- ", b.v)
					if i == b.v {
						b = b.next
					}
				}
				b = b.next
			}
		}
	}
	fmt.Println("}")
}
