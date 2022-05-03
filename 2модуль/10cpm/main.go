package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	s "strings"
)

var (
	keys   map[string]int
	g      [10000000]Point
	n      int
	a      [10000000]Elem
	b      [10000000]Back
	ja, jb int
	i3     int
	help   [5000]*Point
)

type Point struct {
	next      *Elem
	back      *Back
	time      int
	x         string
	dist      int
	ifrom     int
	used      bool
	visited   bool
	blue, red bool
	from      []*Point
}

type Elem struct {
	next *Elem
	v    *Point
}

type Back struct {
	back *Back
	v    *Point
}

func isPresent(u *Point, v *Point) bool {
	x := u.next
	for x != nil {
		if v == x.v {
			return false
		}
		x = x.next
	}
	return true
}

func find_blue(v *Point) {
	v.used = true
	x := v.next
	for x != nil {
		u := x.v
		if !u.visited && u.used || u == v {
			u.blue = true
		}

		if !u.used {
			find_blue(u)
		}
		x = x.next
	}
	v.visited = true
}

func color_blue(v *Point) {
	v.blue = true
	x := v.next
	for x != nil {
		u := x.v
		if !u.blue {
			color_blue(u)
		}
		x = x.next
	}
}

func color_red(v *Point) {
	v.red = true
	for i := 0; i < v.ifrom; i++ {
		from := v.from[i]
		if !from.red {
			color_red(from)
		}
	}
}

func Sort(v *Point) {
	v.used = true
	x := v.next
	for x != nil {
		u := x.v
		if !u.used && !u.blue {
			Sort(u)
		}
		x = x.next
	}
	if !v.blue {
		v.visited = true
		help[i3] = v
		i3++
	}
}

func main() {
	var u int
	bytes, _ := ioutil.ReadAll(os.Stdin)
	text := string(bytes)
	text = s.ReplaceAll(text, " ", "")
	text = s.ReplaceAll(text, "\n", "")
	str := s.Split(text, ";")
	keys = make(map[string]int)

	for i1 := 0; i1 < len(str); i1++ {
		strs := s.Split(str[i1], "<")
		for i := 0; i < len(strs); i++ {
			x := strs[i]
			nx := len(x)
			var v int
			if x[nx-1] == ')' {
				j := nx - 1
				for x[j] != '(' {
					j--
				}
				ts := x[j+1 : nx-1]
				time, _ := strconv.Atoi(ts)
				g[n].time = time
				x = x[:j]
				g[n].x = x
				keys[x] = n
				v = n
				n++
			} else {
				v = keys[x]
			}
			if i != 0 && isPresent(&g[u], &g[v]) {

				if g[v].back == nil {
					b[jb].v = &g[u]
					g[v].back = &b[jb]
				} else {
					b[jb].v = &g[u]
					t1 := g[v].back
					g[v].back = &b[jb]
					b[jb].back = t1
				}
				if g[u].next == nil {
					a[ja].v = &g[v]
					g[u].next = &a[ja]
				} else {
					a[ja].v = &g[v]
					t := g[u].next
					g[u].next = &a[ja]
					a[ja].next = t
				}
				ja++
				jb++
			}
			u = v
		}
	}

	for i := 0; i < n; i++ {
		x := g[i].next
		for x != nil {
			x = x.next
		}
	}

	for i := 0; i < n; i++ {
		if !g[i].used {
			find_blue(&g[i])
		}
	}

	for i := 0; i < n; i++ {
		if g[i].blue {
			color_blue(&g[i])
		}
	}
	for i := 0; i < n; i++ {
		g[i].used = false
		g[i].visited = false
		g[i].dist = g[i].time
		g[i].from = make([]*Point, 50)
	}
	for i := 0; i < n; i++ {
		if !g[i].used {
			Sort(&g[i])
		}
	}

	for i := 0; i*2 < i3; i++ {
		help[i], help[i3-i-1] = help[i3-i-1], help[i]
	}

	max_dist := 0
	var max_v [50000]*Point
	for i := 0; i < i3; i++ {
		v := help[i]
		dist := 0
		x := v.back
		for x != nil {
			from := x.v
			if from.dist > dist {
				dist = from.dist
			}
			x = x.back
		}
		x1 := v.back
		for x1 != nil {
			from := x1.v
			if from.dist == dist {
				v.from[v.ifrom] = from
				v.ifrom++
			}
			x1 = x1.back
		}
		v.dist += dist
		if v.dist > max_dist {
			max_dist = v.dist
		}
	}
	imax := 0
	for i := 0; i < i3; i++ {
		v := help[i]
		if v.dist == max_dist {
			max_v[imax] = v
			imax++
		}
	}

	for i := 0; i < imax; i++ {
		color_red(max_v[i])
	}

	fmt.Println("digraph {")
	for i := 0; i < n; i++ {
		fmt.Printf("%s [label = \"%s(%d)\"", g[i].x, g[i].x, g[i].time)
		if g[i].blue {
			fmt.Printf(", color = blue ")
		}
		if g[i].red {
			fmt.Printf(", color = red ")
		}
		fmt.Printf("]\n")
	}
	for i := 0; i < n; i++ {
		x := g[i].next
		for x != nil {
			fmt.Printf("%s -> %s", g[i].x, x.v.x)
			if g[i].blue {
				fmt.Printf(" [color = blue]")
			}
			find := false
			for _, from := range x.v.from {
				if from == &g[i] {
					find = true
				}
			}
			if g[i].red && x.v.red && find {
				fmt.Printf(" [color = red]")
			}
			fmt.Printf("\n")
			x = x.next
		}
	}
	fmt.Println("}")
}
