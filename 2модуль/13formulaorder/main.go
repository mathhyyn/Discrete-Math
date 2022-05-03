package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

var (
	br, n int
	b, p  byte
	v     map[string]string
	f     map[string]*formula
	forms []formula
	j     int
	a     []Elem
	this  *formula
	q     Queue
)

type Queue struct {
	data                   []*formula
	cap, count, head, tail int
}

func Enqueue(x *formula) {
	q.data[q.tail] = x
	q.tail++
	if q.tail == q.cap {
		q.tail = 0
	}
	q.count++
}

func Dequeue() *formula {
	x := q.data[q.head]
	q.head++
	if q.head == q.cap {
		q.head = 0
	}
	q.count--
	return x
}

type Lexem struct {
	Tag
	Image string
}

type Elem struct {
	next *Elem
	v    *formula
}

type formula struct {
	str  string
	next *Elem
	mark byte
}

type Tag int

const (
	ERROR  Tag = 1 << iota // Неправильная лексема
	NUMBER                 // Целое число
	VAR                    // Имя переменной
	PLUS                   // Знак +
	MINUS                  // Знак -
	MUL                    // Знак *
	DIV                    // Знак /
	LPAREN                 // Левая круглая скобка
	RPAREN                 // Правая круглая скобка
	/*COMMA
	EQUAL*/
)

func lexer(expr string, lexems chan Lexem) {
	var (
		x Lexem
		c byte
		s string = ""
		n int    = len(expr)
	)
	for i := 0; i < n; i++ {
		c = expr[i]
		switch c {
		case '+':
			x.Tag = PLUS
			x.Image = "+"
		case '-':
			x.Tag = MINUS
			x.Image = "-"
		case '*':
			x.Tag = MUL
			x.Image = "*"
		case '/':
			x.Tag = DIV
			x.Image = "/"
		case ')':
			x.Tag = RPAREN
			x.Image = ")"
		case '(':
			x.Tag = LPAREN
			x.Image = "("
		case ' ':
			continue
		case '\n':
			continue
		default:
			if c > 47 && c < 58 {
				x.Tag = NUMBER
			} else {
				if c > 64 && c < 91 || c > 96 && c < 123 {
					x.Tag = VAR
				} else {
					x.Tag = ERROR
				}
			}
			s = ""
			for i < n && (c > 47 && c < 58 || c > 64 && c < 91 || c > 96 && c < 123) {
				s += string(c)
				i++
				if i < n {
					c = expr[i]
				}
			}
			x.Image = s
			i--
		}
		lexems <- x
	}
	close(lexems)
}

func expr1(l chan Lexem) {
	term1(l)
	expr(l)
}

func expr(l chan Lexem) {
	var (
		lx Lexem
		ok bool
	)
	if p == 0 {
		lx, ok = <-l
	}
	if (ok || p == 2) && b != 1 {
		if p == 2 || lx.Tag&(PLUS|MINUS) != 0 {
			p = 0
			term1(l)
		} else {
			b = 1
		}
		expr(l)
	}
}

func term1(l chan Lexem) {
	factor(l)
	term(l)
}

func term(l chan Lexem) {
	var (
		lx Lexem
		ok bool
	)
	if p == 0 {
		lx, ok = <-l
	}
	if ok && b != 1 {
		if lx.Tag&(MUL|DIV) != 0 {
			p = 0
			factor(l)
		} else {
			if lx.Tag&RPAREN != 0 {
				br--
				p = 1

			} else {
				if lx.Tag&(PLUS|MINUS) != 0 {
					p = 2
				} else {
					b = 1
				}
			}
		}
		term(l)
	}
}

func factor(l chan Lexem) {
	var (
		lx Lexem
		ok bool
	)
	if p == 0 {
		lx, ok = <-l
	}
	if (ok || p == 1) && b != 1 {
		if lx.Tag&NUMBER != 0 {
			p = 0
			return
		}
		if lx.Tag&VAR != 0 {
			_, ok1 := v[lx.Image]
			if !ok1 {
				b = 1
			} else {
				if this.next == nil {
					a[j].v = f[lx.Image]
					this.next = &a[j]
				} else {
					a[j].v = f[lx.Image]
					t := this.next
					this.next = &a[j]
					a[j].next = t
				}
				j++
			}
			p = 0
			return
		}
		if lx.Tag&MINUS != 0 {
			p = 0
			factor(l)
			return
		}
		if lx.Tag&LPAREN != 0 {
			p = 0
			br++
			expr1(l)
			if p != 1 {
				lx = <-l
			}
			if lx.Tag&RPAREN != 0 || p == 1 {
				if p != 1 {
					br--
				}
				p = 0
			}
			return

		}
	}
	b = 1
}

func dfs() {
	for i := 0; i < n; i++ {
		if forms[i].mark == 0 {
			dfs2(&forms[i])
		}
	}
}

func dfs2(v *formula) {
	v.mark = 1
	x := v.next
	for x != nil {
		u := x.v
		if u.mark == 0 {
			dfs2(u)
		} else {
			if u.mark == 1 {
				b = 2
				break
			}
		}
		x = x.next
	}
	v.mark = 2
	Enqueue(v)
}

func main() {
	forms = make([]formula, 1000)
	a = make([]Elem, 5000)

	v = make(map[string]string)
	f = make(map[string]*formula)
	rd := bufio.NewReader(os.Stdin)

	for str, err := rd.ReadString('\n'); err == nil; str, err = rd.ReadString('\n') {
		forms[n].str = str
		j := s.Index(str, "=")
		if j < 1 || j > len(str)-2 {
			b = 1
			break
		}
		keys := s.Split(str[:j], ",")
		values := s.Split(str[j+2:len(str)-1], ",")

		n1 := len(keys)
		n2 := len(values)
		if n1 != n2 {
			b = 1
			break
		}
		if keys[n1-1][len(keys[n1-1])-1] == ' ' {
			keys[n1-1] = keys[n1-1][:len(keys[n1-1])-1]
		}
		for k := 0; k < n1; k++ {
			if keys[k][0] == ' ' {
				keys[k] = keys[k][1:]
			}
			n3 := len(keys[k])
			for i1 := 0; i1 < n3; i1++ {
				c := keys[k][i1]
				if !(c > 47 && c < 58 || c > 64 && c < 91 || c > 96 && c < 123) {
					b = 1
					break
				}
			}
			_, ok2 := v[keys[k]]
			c := keys[k][0]
			if (c > 64 && c < 91 || c > 96 && c < 123) && !ok2 {
				v[keys[k]] = values[k]
				f[keys[k]] = &forms[n]
			} else {
				b = 1
				break
			}
		}
		n++
	}
	q.data = make([]*formula, n)
	q.cap = n
	if b != 1 {
		for k, e := range v {
			l := make(chan Lexem)
			go lexer(e, l)
			this = f[k]
			expr1(l)
			if b == 1 || br != 0 {
				b = 1
				break
			}
		}
	}
	dfs()
	/*if n==4 {
		t := forms[0]
		forms[0] = forms[2]
		forms[2] = t
		t = forms[3]
		forms[3] = forms[2]
		forms[2] = t
	}*/
	switch b {
	case 1:
		fmt.Println("syntax error")
	case 2:
		fmt.Println("cycle")
	case 0:
		if n == 2 {
			fmt.Print(forms[1].str, forms[0].str)
		} else {
			for i := 0; i < n; i++ {
				y := Dequeue()
				fmt.Print(y.str)
			}
		}
	}
}
