package main

import (
	"fmt"
)

var (
	n, N, j1 int
	key      map[int]*Point
	Loops    []*Point
	g        []Point
	a        []Elem
	b        []Back
)

type Elem struct {
	next *Elem
	v    *Point
}

type Back struct {
	back *Back
	v    *Point
}

type Point struct {
	x, o     int
	s        string
	next     *Elem
	back     *Back
	used     bool
	parent   *Point
	sdom     *Point
	label    *Point
	ancestor *Point
	bucket   []*Point
	t        int
	dom      *Point
}

/*type stack struct {
	data     []*Point
	cap, top int
}

func InitStack(s *stack) {
	s.data = make([]*Point, N)
	s.cap = N
}

func Push(s *stack, x *Point) {
	s.data[s.top] = x
	s.top++
}

func Pop(s *stack) *Point {
	s.top--
	return s.data[s.top]
}*/

func dfs(v *Point) {
	v.t = N
	Loops[N] = v
	//fmt.Println(v.x)
	N++
	v.used = true
	x := v.next
	for x != nil {
		u := x.v
		if u.back == nil {
			b[j1].v = v
			u.back = &b[j1]
		} else {
			b[j1].v = v
			t1 := u.back
			u.back = &b[j1]
			b[j1].back = t1
		}
		j1++
		if !u.used {
			u.parent = v
			dfs(u)
		}
		x = x.next
	}

}

func SearchAndCut(v *Point) *Point {
	if v.ancestor == nil {
		return v
	} else {
		root := SearchAndCut(v.ancestor)
		if v.ancestor.label.sdom.t < v.label.sdom.t {
			v.label = v.ancestor.label
		}
		v.ancestor = root
		return root
	}
}

func FindMin(v *Point) *Point {
	SearchAndCut(v)
	return v.label
}

func Dominators() {
	for i := N - 1; i > 0; i-- {
		w := Loops[i]
		x := w.back
		for x != nil {
			v := x.v
			u := FindMin(v)
			if u.sdom.t < w.sdom.t {
				w.sdom = u.sdom
			}
			x = x.back
		}
		w.ancestor = w.parent
		w.sdom.bucket = append(w.sdom.bucket, w)
		for _, v := range w.parent.bucket {
			u := FindMin(v)
			if u.sdom == v.sdom {
				v.dom = v.sdom
			} else {
				v.dom = u
			}
		}
		w.parent.bucket = nil
	}
	for i := 1; i < N; i++ {
		w := Loops[i]
		if w.dom != w.sdom {
			w.dom = w.dom.dom
		}
	}
	g[0].dom = nil
}

func FindLoops() int {
	l := 0
	for i := 0; i < N; i++ {
		v := Loops[i]
		x := v.back
		for x != nil {
			u := x.v
			for u != v && u != nil {
				u = u.dom
			}
			if u == v {
				l++
				break
			}
			x = x.back
		}
	}
	return l
}

func main() {
	var (
		x, o int
		s    string
	)
	fmt.Scan(&n)
	g = make([]Point, n)
	Loops = make([]*Point, n)
	a = make([]Elem, 2*n)
	b = make([]Back, 2*n)
	j := 0
	key = make(map[int]*Point, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		g[i].x = x
		fmt.Scan(&s)
		g[i].s = s
		key[x] = &g[i]
		if s[0] != 'A' {
			fmt.Scan(&o)
			g[i].o = o
		} else {
			g[i].o = -1
		}
	}
	for i := 0; i < n; i++ {
		g[i].sdom = &g[i]
		g[i].label = &g[i]
		switch g[i].s[0] {
		case 'A':
			if i < n-1 {
				a[j].v = &g[i+1]
				g[i].next = &a[j]
				/*if g[i+1].back == nil {
					b[j1].v = &g[i]
					g[i+1].back = &b[j1]
				} else {
					b[j1].v = &g[i]
					t1 := g[i+1].back
					g[i+1].back = &b[j1]
					b[j1].back = t1
				}*/
			}
		case 'B':
			a[j].v = key[g[i].o]
			g[i].next = &a[j]
			/*if key[g[i].o].back == nil {
				b[j1].v = &g[i]
				key[g[i].o].back = &b[j1]
			} else {
				b[j1].v = &g[i]
				t1 := key[g[i].o].back
				key[g[i].o].back = &b[j1]
				b[j1].back = t1
			}*/
			if i < n-1 {
				j++
				a[j].v = &g[i+1]
				t := g[i].next
				g[i].next = &a[j]
				a[j].next = t
				/*j1++
				b[j1].v = &g[i]
				t1 := g[i+1].back
				g[i+1].back = &b[j1]
				b[j1].back = t1*/
			}
		case 'J':
			a[j].v = key[g[i].o]
			g[i].next = &a[j]
			/*if key[g[i].o].back == nil {
				b[j1].v = &g[i]
				key[g[i].o].back = &b[j1]
			} else {
				b[j1].v = &g[i]
				t1 := key[g[i].o].back
				key[g[i].o].back = &b[j1]
				b[j1].back = t1
			}*/
		}
		j++
		//j1++
	}
	dfs(&g[0])
	Dominators()
	l := 0
	for i := 0; i < N; i++ {
		v := Loops[i]
		x := v.back
		for x != nil {
			u := x.v
			for u != v && u != nil {
				u = u.dom
			}
			if u == v {
				l++
				break
			}
			x = x.back
		}
	}
	fmt.Println(l)
}
