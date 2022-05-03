package main

import "fmt"

var (
	n, m, q, order int
	A, A1          []cond
)

type cond struct {
	T    []int
	F    []string
	used bool
	new  int
}

func dfs(i int) {
	A[i].used = true
	A[i].new = order
	order++
	for j := 0; j < m; j++ {
		if !A[A[i].T[j]].used {
			dfs(A[i].T[j])
		}
	}
}

func main() {
	var (
		x int
		s string
	)
	fmt.Scan(&n, &m, &q)
	A = make([]cond, n)
	A1 = make([]cond, n)
	for i := 0; i < n; i++ {
		A[i].T = make([]int, m)
		A1[i].T = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&x)
			A[i].T[j] = x
		}
	}
	for i := 0; i < n; i++ {
		A[i].F = make([]string, m)
		A1[i].F = make([]string, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&s)
			A[i].F[j] = s
		}
	}
	dfs(q)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			A1[A[i].new].T[j] = A[A[i].T[j]].new
			A1[A[i].new].F[j] = A[i].F[j]
		}
	}

	fmt.Print(order, "\n", m, "\n", 0, "\n")
	for i := 0; i < order; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(A1[i].T[j], " ")
		}
		fmt.Println()
	}
	for i := 0; i < order; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(A1[i].F[j], " ")
		}
		fmt.Println()
	}
}
