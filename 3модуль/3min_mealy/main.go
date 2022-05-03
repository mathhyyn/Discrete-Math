package main

import "fmt"

type cond struct {
	i, depth int
	parent   *cond
}
type cond1 struct {
	used      bool
	ispresent bool
	new       int
}

var (
	N, M, q0  int
	order, N1 int
	D, D1     [][]int
	F, F1     [][]string
	Q, pi     []cond
	A         []cond1
)

func dfs(i int) {
	A[i].used = true
	A[i].new = order
	order++
	for j := 0; j < M; j++ {
		if !A[D1[i][j]].used && A[D1[i][j]].ispresent {
			dfs(D1[i][j])
		}
	}
}

func Find(x *cond) *cond {
	if x.parent == x {
		return x
	} else {
		y := Find(x.parent)
		x.parent = y
		return y
	}
}

func Union(x, y *cond) {
	if x.depth < y.depth {
		x.parent = y
	} else {
		y.parent = x
		if x.depth == y.depth && x != y {
			x.depth++
		}
	}
}

func Split1() int {
	m := N
	for i := 0; i < N; i++ {
		Q[i].parent = &Q[i]
	}

	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			q1 := Find(&Q[i])
			q2 := Find(&Q[j])
			if q1 != q2 {
				eq := true
				for x := 0; x < M; x++ {
					q1_ := Q[i].i
					q2_ := Q[j].i
					if F[q1_][x] != F[q2_][x] {
						eq = false
						break
					}
				}
				if eq {
					Union(q1, q2)
					m--
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		q := &Q[i]
		pi[q.i] = *Find(q)
	}
	return m
}

func Split() int {
	m1 := N
	for i := 0; i < N; i++ {
		Q[i].parent = &Q[i]
	}

	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			q1 := &Q[i]
			q2 := &Q[j]
			q1_ := Q[i].i
			q2_ := Q[j].i
			if pi[q1_] == pi[q2_] && Find(q1) != Find(q2) {
				eq := true
				for x := 0; x < M; x++ {
					w1 := D[q1_][x]
					w2 := D[q2_][x]
					if pi[w1] != pi[w2] {
						eq = false
						break
					}
				}
				if eq {
					Union(q1, q2)
					m1--
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		q := &Q[i]
		pi[q.i] = *Find(q)
	}
	return m1
}

func AufenkampHohn() {
	m := Split1()
	for {
		m1 := Split()
		if m == m1 {
			break
		}
		m = m1
	}

	for i := 0; i < N; i++ {
		q := Q[i].i
		q1 := pi[q]
		if !A[q1.i].ispresent {
			A[q1.i].ispresent = true
			N1++
			for x := 0; x < M; x++ {
				D1[q1.i][x] = pi[D[q][x]].i
				F1[q1.i][x] = F[q][x]
			}
			q0 = pi[q0].i
		}
	}
}

func main() {
	fmt.Println("digraph {\nrankdir = LR\ndummy [label = \"\", shape = none]")

	fmt.Scan(&N, &M, &q0)
	D, D1 = make([][]int, N), make([][]int, N)
	F, F1 = make([][]string, N), make([][]string, N)
	Q, pi = make([]cond, N), make([]cond, N)
	A = make([]cond1, N)

	for i := 0; i < N; i++ {
		A[i].new = -1
	}

	for i := 0; i < N; i++ {
		Q[i].i = i
		D[i], D1[i] = make([]int, M), make([]int, M)
		for j := 0; j < M; j++ {
			fmt.Scan(&D[i][j])
		}
	}
	for i := 0; i < N; i++ {
		F[i], F1[i] = make([]string, M), make([]string, M)
		for j := 0; j < M; j++ {
			fmt.Scan(&F[i][j])
		}
	}

	AufenkampHohn()

	dfs(q0)

	for i := 0; i < N; i++ {
		if A[i].ispresent && A[i].new != -1 {
			fmt.Println(A[i].new, "[shape = circle]")
		}
	}
	fmt.Println("dummy ->", 0)
	for i := 0; i < N; i++ {
		if A[i].ispresent && A[i].new != -1 {
			for j := 0; j < M; j++ {
				if A[D1[i][j]].ispresent {
					fmt.Print(A[i].new, " -> ", A[D1[i][j]].new, " [label = \"")
					fmt.Printf("%c(%s)\"]\n", 97+j, F1[i][j])
				}
			}
		}
	}
	fmt.Println("}")
}
