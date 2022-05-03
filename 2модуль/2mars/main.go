package main

import (
	"fmt"
)

var (
	m []crew
	n int
)

type crew struct {
	a           []byte
	i, k, g, g1 int
	used        bool
}

type team struct {
	a []byte
	k int
}

func comp(t team, i int) bool {
	for j := 0; j < n; j++ {
		if t.a[j] == 1 {
			if m[i].a[j] == '+' {
				return false
			}
		}
	}
	return true
}

func find(j, k int) int {
	h := m[j].i
	for i := k; i < n; i++ {
		if m[i].a[h] == '+' {
			return i
		}
	}
	return 0
}

func next(i, k, l, g, g1 int, f bool) (int, int) {
	h := find(i, k)
	if h != 0 && !m[h].used {
		m[h].used = true
		if f {
			g++
		} else {
			g1++
		}
		//fmt.Println("next", h, m[h].i+1, f)
		loc := h - 1
		for loc > i+l {
			t := m[loc+1]
			m[loc+1] = m[loc]
			m[loc] = t
			loc--
		}
		g, g1 = next(i+l+1, i+l+2, 0, g, g1, !f)
		g, g1 = next(i, h+1, l+1, g, g1, f)
	}
	return g, g1
}

func main() {
	fmt.Scan(&n)
	var (
		s      string
		t1, t2 team
		k, f   byte
	)
	m = make([]crew, n)
	t1.a = make([]byte, n)
	t2.a = make([]byte, n)

	for i := 0; i < n; i++ {
		m[i].a = make([]byte, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&s)
			m[i].a[j] = s[0]
			m[i].i = i
			if m[i].a[j] == '+' {
				m[i].k++
			}
		}
		fmt.Scanln()
	}

	for i := 0; i < n; i++ {
		if comp(t2, i) {
			k = 2
		}
		if comp(t1, i) {
			k++
		}
		if k == 0 {
			f = 1
			break
		}
		g, g1 := next(i, i+1, 0, 0, 0, true)
		//fmt.Println(i, m[i].i+1, t1.k, g1+1, t2.k, g)
		if (k == 3 && g1+1 == g) || k == 1 {
			t1.a[m[i].i] = 1
			t1.k++
			m[i].g = -1
		} else {
			if k == 3 {
				m[i].g = g
				m[i].g1 = g1
				i += g + g1
			} else {
				t2.a[m[i].i] = 1
				t2.k++
				m[i].g = -1
			}
		}
		k = 0
	}

	for i := 0; i < n; i++ {
		if m[i].g != -1 {
			if comp(t2, i) {
				k = 2
			}
			if comp(t1, i) {
				k++
			}
			if k == 0 {
				f = 1
				break
			}
			//fmt.Println(i, m[i].i+1, t1.k, m[i].g1+1)
			if (k == 3 && t1.k+m[i].g1+1 <= n/2) || k == 1 {
				t1.a[m[i].i] = 1
				t1.k++
			} else {
				t2.a[m[i].i] = 1
				t2.k++
			}
			k = 0
		}
	}

	if f == 1 {
		fmt.Println("No solution")
	} else {
		for i := 0; i < n; i++ {
			if t1.a[i] == 1 {
				fmt.Print(i+1, " ")
			}
		}
	}
}
