package main

import "fmt"

var (
	time, count int = 1, 1
	G           []point
	n, m        int
	s           stack
)

type condensation struct {
	isbase bool
	min    int
}

type point struct {
	v, T1, comp, low int
	next             *Elem
}

type Elem struct {
	next *Elem
	v    int
}

type stack struct {
	data     []int
	cap, top int
}

func InitStack() {
	s.data = make([]int, n)
	s.cap = n
}

func Push(x int) {
	s.data[s.top] = x
	s.top++
}

func Pop() int {
	s.top--
	return s.data[s.top]
}

func Tarjan() {
	InitStack()
	for v := 0; v < n; v++ {
		if G[v].T1 == 0 {
			VisitVertex_Tarjan(v)
		}
	}
}

func VisitVertex_Tarjan(v int) {
	G[v].T1, G[v].low = time, time
	time++
	Push(v)
	x := G[v].next
	for x != nil {
		if G[x.v].T1 == 0 {
			VisitVertex_Tarjan(x.v)
		}
		if G[x.v].comp == 0 && G[v].low > G[x.v].low {
			G[v].low = G[x.v].low
		}
		x = x.next
	}
	if G[v].T1 == G[v].low {
		for {
			u := Pop()
			G[u].comp = count
			if u == v {
				break
			}
		}
		count++
	}
}

func main() {
	var x, y int
	fmt.Scan(&n)
	fmt.Scan(&m)
	G = make([]point, n)
	a := make([]Elem, m)
	j := 0
	for i := 0; i < m; i++ {
		fmt.Scan(&x, &y)
		if G[x].next == nil {
			a[j].v = y
			G[x].next = &a[j]
		} else {
			a[j].v = y
			t := G[x].next
			G[x].next = &a[j]
			a[j].next = t
		}
		j++
	}

	Tarjan()

	condensation := make([]condensation, count)

	for i := n - 1; i >= 0; i-- {
		condensation[G[i].comp].min = i
		x := G[i].next
		a := G[i].comp
		for x != nil {
			b := G[x.v].comp
			if a != b {
				condensation[b].isbase = true
			}
			x = x.next
		}
	}

	for i := 1; i < count; i++ {
		if !condensation[i].isbase {
			fmt.Println(condensation[i].min)
		}
	}
}
