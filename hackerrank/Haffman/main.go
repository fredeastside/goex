package main

import (
	"fmt"
	"sort"
	"strings"
)

var keyLength int
var fullCode string

type queue []*tree

type tree struct {
	left  *tree
	right *tree
	key   string
	code  string
	count int
}

func (q queue) Len() int {
	return len(q)
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q queue) Less(i, j int) bool {
	return q[i].count < q[j].count
}

func (t tree) String() string {
	return fmt.Sprintf("{key: '%s', count: '%d', code: '%s'}\n", t.key, t.count, t.code)
}

func walk(tr *tree, code string, print bool) {
	if tr == nil {
		return
	}
	if tr.key != "" {
		tr.code = code
		if print {
			fmt.Printf("%s: %s\n", tr.key, tr.code)
		}
	}
	walk(tr.left, code+"0", print)
	walk(tr.right, code+"1", print)
}

func findCode(tr *tree, key string) string {
	if tr == nil {
		return ""
	}

	if tr.key == key {
		return tr.code
	}

	code := findCode(tr.left, key)
	if code != "" {
		return code
	}

	return findCode(tr.right, key)
}

func main() {
	//str := "beep boop beer!"
	//str := "abacabad"
	str := "bcbdbcbebcbdbcb"
	//str := "vjdruvnkxnvjkfnvjxdfgivdnvjbxknvjdinvndkjsnjnsbfjkdnbjntubnknjksnjbnbndvjvspmlzaqwqjhdh"
	str = strings.TrimSpace(str)
	var q queue

	for _, key := range str {
		tr := tree{key: string(key), count: 1}
		found := false
		for _, existTr := range q {
			if tr.key == existTr.key {
				existTr.count++
				found = true
				break
			}
		}
		if !found {
			keyLength++
			q = append(q, &tr)
		}
	}

	if len(q) == 1 {
		fmt.Printf("%d %d\n", 1, 1)
		fmt.Printf("%s: %s\n", q[0].key, "0")
		fmt.Printf("%s\n", "0")

		return
	}
	//fmt.Printf("%+v", q)

	for {
		if len(q) == 1 {
			break
		}
		sort.Sort(q)
		tr1 := q[0]
		tr2 := q[1]
		q = append(q[:0], q[2:]...)
		tr := tree{left: tr1, right: tr2, count: tr1.count + tr2.count}
		q = append(q, &tr)
	}

	walk(q[0], "", false)
	for _, key := range str {
		code := findCode(q[0], string(key))
		fullCode += code
	}
	fmt.Printf("%d %d\n", keyLength, len(fullCode))
	walk(q[0], "", true)
	fmt.Printf("%s\n", fullCode)
}
