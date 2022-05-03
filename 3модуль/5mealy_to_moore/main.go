package main

import "fmt"

type cond struct {
	i int
	y string
}

type pair struct {
	q cond
	x string
}

type mealy struct {
	T []int
	F []string
}

func main() {
	fmt.Println("digraph {\nrankdir = LR")
	var (
		A    []mealy
		K, m int
		n, N int
		X, Y []string
		Q    []cond
		Q1   map[cond]int
		D    map[pair]cond
		q    cond
		p    pair
	)

	fmt.Scan(&K)
	X = make([]string, K)
	for i := 0; i < K; i++ {
		fmt.Scan(&X[i])
	}

	D = make(map[pair]cond)
	Q1 = make(map[cond]int)

	fmt.Scan(&m)
	Y = make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&Y[i])
	}

	fmt.Scan(&n)
	A = make([]mealy, n)
	for i := 0; i < n; i++ {
		A[i].T = make([]int, K)
		for j := 0; j < K; j++ {
			x := 0
			fmt.Scan(&x)
			A[i].T[j] = x

		}
	}
	for i := 0; i < n; i++ {
		A[i].F = make([]string, K)
		for j := 0; j < K; j++ {
			x := ""
			fmt.Scan(&x)
			A[i].F[j] = x
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			t := 0
			for k := 0; k < n; k++ {
				for l := 0; l < K; l++ {
					if A[k].F[l] == Y[j] && A[k].T[l] == i {
						q.i = i
						q.y = Y[j]
						Q = append(Q, q)
						N++
						t = 1
						break
					}
				}
				if t == 1 {
					break
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		t := Q[i]
		Q1[t] = i
		for j := 0; j < K; j++ {
			q.i = A[t.i].T[j]
			q.y = A[t.i].F[j]
			p.q = t
			p.x = X[j]
			D[p] = q
		}
	}

	for i := 0; i < N; i++ {
		fmt.Print(i, " [label = \"(", Q[i].i, ",", Q[i].y, ")\"]\n")
	}
	for i := 0; i < N; i++ {
		p.q = Q[i]
		for j := 0; j < K; j++ {
			p.x = X[j]
			q, ok := D[p]
			if ok {
				fmt.Print(i, " -> ", Q1[q], " [label = \"", p.x, "\"]\n")
			}
		}
	}
	fmt.Println("}")
}
