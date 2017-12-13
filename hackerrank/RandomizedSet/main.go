package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	arr       []int
	hashTable map[int]int
}

func NewRandomizedSet() *RandomizedSet {
	rs := RandomizedSet{hashTable: make(map[int]int)}

	return &rs
}

func (r *RandomizedSet) Insert(value int) bool {
	if r.Contains(value) {
		return false
	}

	key := len(r.arr)
	r.hashTable[value] = key
	r.arr = append(r.arr, value)

	return true
}

func (r *RandomizedSet) Contains(value int) bool {
	_, ok := r.hashTable[value]

	return ok
}

func (r *RandomizedSet) Remove(value int) bool {
	if r.Contains(value) {
		key := r.hashTable[value]
		r.arr[key] = r.arr[len(r.arr)-1]
		r.arr = r.arr[:len(r.arr)-1]
		delete(r.hashTable, value)

		return true
	}

	return false
}

func (r *RandomizedSet) Random() int {
	s := rand.NewSource(time.Now().Unix())
	rand := rand.New(s)

	return r.arr[rand.Intn(len(r.arr))]
}

func main() {
	rs := NewRandomizedSet()
	fmt.Println(rs.Insert(1))
	fmt.Println(rs.Remove(2))
	rs.Insert(2)
	fmt.Println(rs)
	fmt.Println(rs.Random())
	fmt.Println(rs.Remove(1))
	fmt.Println(rs.Insert(2))
	fmt.Println(rs.Random())
}
