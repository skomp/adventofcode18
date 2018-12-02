package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input")
	s := bufio.NewScanner(f)
	c := []string{}
	for s.Scan() {
		c = append(c, s.Text())
	}
	one(c)
	two(c)
}

func one(c []string) {
	c2, c3 := 0, 0
	for _, l := range c {
		lstat := map[rune]int{}
		for _, r := range l {
			lstat[r]++
		}
		i, j := 0, 0
		for _, r := range lstat {
			if r == 3 {
				i = 1
			}
			if r == 2 {
				j = 1
			}

		}
		c2 += j
		c3 += i
	}
	println(c2 * c3)
}

func bDist(a, b string) (int, int) {
	d := 0
	rp := -1
	for p := range a {
		if len(b) < p {
			break
		}
		if a[p] != b[p] {
			d++
			rp = p
		}
	}
	return d, rp
}

func two(c []string) {
	for p, a := range c {
		//if p == len(a)-1 {
		//	break
		//}
		for _, b := range c[p+1:] {
			d, rp := bDist(a, b)
			if d == 1 {
				fmt.Printf("%s%s", a[:rp], a[rp+1:])
				return
			}
		}
	}
}
