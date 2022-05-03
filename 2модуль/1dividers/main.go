package main

import (
	"fmt"
	"math"
)

var d []int
var n, q int

func dividers() {
	a := 1
	q = 0
	for a*a <= n {
		d[q] = a
		q++
		if a*a != n {
			d[q] = n / a
			q++
		}
		a++
		for a*a < n && n%a != 0 {
			a++
		}
	}
}

func main() {
	var d0 [1000000]int
	var t, i int
	fmt.Scan(&n)
	n1 := int(math.Sqrt(float64(n))) * 2
	d = d0[:n1]
	dividers()
	for i = 0; i < q; i++ {
		loc := i - 1
		for loc >= 0 && d[loc+1] < d[loc] {
			t := d[loc+1]
			d[loc+1] = d[loc]
			d[loc] = t
			loc--
		}
	}
	fmt.Println("graph {")
	for i = q - 1; i >= 0; i-- {
		fmt.Println(d[i])
	}

	for i = q - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if d[i]%d[j] == 0 {
				for k := i - 1; k > j; k-- {
					if d[i]%d[k] == 0 && d[k]%d[j] == 0 {
						t = 1
						break
					}
				}
				if t != 1 {
					fmt.Println(d[i], "--", d[j])
				}
				t = 0
			}
		}

	}
	fmt.Println("}")
}
